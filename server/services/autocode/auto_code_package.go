package autocode

import (
	"context"
	"fmt"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/kesilent/react-light-blog/dal/request"
	"github.com/kesilent/react-light-blog/global"
	"github.com/kesilent/react-light-blog/utils/ast"
	"github.com/pkg/errors"
)

var AutoCodePackage = new(autoCodePackage)

type autoCodePackage struct{}

// Create 创建包信息
func (s *autoCodePackage) Create(ctx context.Context, info *request.AutoCodeReq) error {
	switch {
	case info.Template == "":
		return errors.New("模板不能为空!")
	case info.Template == "page":
		return errors.New("page为表单生成器!")
	case info.PackageName == "":
		return errors.New("PackageName不能为空!")
	case token.IsKeyword(info.PackageName):
		return errors.Errorf("%s为go的关键字!", info.PackageName)
	case info.Template == "package":
		if info.PackageName == "system" || info.PackageName == "example" {
			return errors.New("不能使用已保留的package name")
		}
	default:
		break
	}

	return nil
}

// 预加载模板文件
func (s *autoCodePackage) Templates(info request.AutoCodeReq, isPackage bool) (code map[string]string, asts map[string]ast.Ast, creates map[string]string, err error) {
	code = make(map[string]string)
	asts = make(map[string]ast.Ast)
	creates = make(map[string]string)

	templateDir := filepath.Join(global.RLB_CONFIG.AutoCode.Root, global.RLB_CONFIG.AutoCode.Server, "resource", info.Template)
	templateDirs, err := os.ReadDir(templateDir)
	if err != nil {
		return nil, nil, nil, errors.Wrapf(err, "读取模版文件夹[%s]失败!", templateDir)
	}

	for i := 0; i < len(templateDirs); i++ {
		second := filepath.Join(templateDir, templateDirs[i].Name())
		switch templateDirs[i].Name() {
		case "server":
			if !info.GenerateServer && !isPackage {
				break
			}
			var secondDirs []os.DirEntry
			secondDirs, err = os.ReadDir(second)
			if err != nil {
				return nil, nil, nil, errors.Wrapf(err, "读取模版文件夹[%s]失败!", second)
			}

			for j := 0; j < len(secondDirs); j++ {
				if secondDirs[j].Name() == ".DS_Store" {
					continue
				}

				three := filepath.Join(second, secondDirs[j].Name())

				switch secondDirs[j].Name() {

				case "api", "router", "service":
					var threeDirs []os.DirEntry
					threeDirs, err = os.ReadDir(three)
					if err != nil {
						return nil, nil, nil, errors.Wrapf(err, "读取模版文件夹[%s]失败!", three)
					}

					for k := 0; k < len(threeDirs); k++ {
						four := filepath.Join(three, threeDirs[k].Name())
						if threeDirs[k].IsDir() {
							return nil, nil, nil, errors.Errorf("[filpath:%s]非法模版文件夹!", four)
						}
						ext := filepath.Ext(four)
						if ext != ".tpl" {
							return nil, nil, nil, errors.Errorf("[filpath:%s]非法模版后缀!", four)
						}

						api := strings.Index(threeDirs[k].Name(), "api")
						hasEnter := strings.Index(threeDirs[k].Name(), "enter")
						router := strings.Index(threeDirs[k].Name(), "router")
						service := strings.Index(threeDirs[k].Name(), "service")
						if router == -1 && api == -1 && service == -1 && hasEnter == -1 {
							return nil, nil, nil, errors.Errorf("[filpath:%s]非法模版文件!", four)
						}
						if hasEnter != -1 {
							isApi := strings.Index(secondDirs[j].Name(), "api")
							isRouter := strings.Index(secondDirs[j].Name(), "router")
							isService := strings.Index(secondDirs[j].Name(), "service")
							if isRouter != -1 {
								pluginRouterEnter := &ast.PluginEnter{
									Type:            ast.TypePluginRouterEnter,
									Path:            filepath.Join(global.RLB_CONFIG.AutoCode.Root, global.RLB_CONFIG.AutoCode.Server, "plugin", info.PackageName, secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext)),
									ImportPath:      fmt.Sprintf(`"%s/plugin/%s/api"`, global.RLB_CONFIG.AutoCode.Module, info.PackageName),
									StructName:      info.StructName,
									StructCamelName: info.Abbreviation,
									ModuleName:      "api" + info.StructName,
									GroupName:       "Api",
									PackageName:     "api",
									ServiceName:     info.StructName,
								}
								asts[pluginRouterEnter.Path+"=>"+pluginRouterEnter.Type.String()] = pluginRouterEnter
								creates[four] = pluginRouterEnter.Path
							}
							if isApi != -1 {
								pluginApiEnter := &ast.PluginEnter{
									Type:            ast.TypePluginApiEnter,
									Path:            filepath.Join(global.RLB_CONFIG.AutoCode.Root, global.RLB_CONFIG.AutoCode.Server, "plugin", info.PackageName, secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext)),
									ImportPath:      fmt.Sprintf(`"%s/plugin/%s/service"`, global.RLB_CONFIG.AutoCode.Module, info.PackageName),
									StructName:      info.StructName,
									StructCamelName: info.Abbreviation,
									ModuleName:      "service" + info.StructName,
									GroupName:       "Service",
									PackageName:     "service",
									ServiceName:     info.StructName,
								}
								asts[pluginApiEnter.Path+"=>"+pluginApiEnter.Type.String()] = pluginApiEnter
								creates[four] = pluginApiEnter.Path
							}
							if isService != -1 {
								pluginServiceEnter := &ast.PluginEnter{
									Type:            ast.TypePluginServiceEnter,
									Path:            filepath.Join(global.RLB_CONFIG.AutoCode.Root, global.RLB_CONFIG.AutoCode.Server, "plugin", info.PackageName, secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext)),
									StructName:      info.StructName,
									StructCamelName: info.Abbreviation,
								}
								asts[pluginServiceEnter.Path+"=>"+pluginServiceEnter.Type.String()] = pluginServiceEnter
								creates[four] = pluginServiceEnter.Path
							}
							continue
						}

						create := filepath.Join(global.RLB_CONFIG.AutoCode.Root, global.RLB_CONFIG.AutoCode.Server, "plugin", info.PackageName, secondDirs[j].Name(), info.HumpPackageName+".go")
						code[four] = create
					}
				case "gen", "config", "initialize", "plugin", "response":
					if info.Template == "package" {
						continue
					} // package模板不需要生成gen, config, initialize
					var threeDirs []os.DirEntry
					threeDirs, err = os.ReadDir(three)
					if err != nil {
						return nil, nil, nil, errors.Wrapf(err, "读取模版文件夹[%s]失败!", three)
					}

					for k := 0; k < len(threeDirs); k++ {
						if threeDirs[k].Name() == ".DS_Store" {
							continue
						}
						four := filepath.Join(three, threeDirs[k].Name())
						if threeDirs[k].IsDir() {
							return nil, nil, nil, errors.Errorf("[filpath:%s]非法模版文件夹!", four)
						}
						ext := filepath.Ext(four)
						if ext != ".tpl" {
							return nil, nil, nil, errors.Errorf("[filpath:%s]非法模版后缀!", four)
						}

						gen := strings.Index(threeDirs[k].Name(), "gen")
						api := strings.Index(threeDirs[k].Name(), "api")
						menu := strings.Index(threeDirs[k].Name(), "menu")
						viper := strings.Index(threeDirs[k].Name(), "viper")
						plugin := strings.Index(threeDirs[k].Name(), "plugin")
						config := strings.Index(threeDirs[k].Name(), "config")
						router := strings.Index(threeDirs[k].Name(), "router")
						hasGorm := strings.Index(threeDirs[k].Name(), "gorm")
						response := strings.Index(threeDirs[k].Name(), "response")
						if gen != -1 && api != -1 && menu != -1 && viper != -1 && plugin != -1 && config != -1 && router != -1 && hasGorm != -1 && response != -1 {
							return nil, nil, nil, errors.Errorf("[filpath:%s]非法模版文件!", four)
						}
						if api != -1 || menu != -1 || viper != -1 || response != -1 || plugin != -1 || config != -1 {
							creates[four] = filepath.Join(global.RLB_CONFIG.AutoCode.Root, global.RLB_CONFIG.AutoCode.Server, "plugin", info.PackageName, secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext))
						}
						if gen != -1 {
							pluginGen := &ast.PluginGen{
								Type:        ast.TypePluginGen,
								Path:        filepath.Join(global.RLB_CONFIG.AutoCode.Root, global.RLB_CONFIG.AutoCode.Server, "plugin", info.PackageName, secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext)),
								ImportPath:  fmt.Sprintf(`"%s/plugin/%s/model"`, global.RLB_CONFIG.AutoCode.Module, info.PackageName),
								StructName:  info.StructName,
								PackageName: "model",
								IsNew:       true,
							}
							asts[pluginGen.Path+"=>"+pluginGen.Type.String()] = pluginGen
							creates[four] = pluginGen.Path
						}
						if hasGorm != -1 {
							pluginInitializeGorm := &ast.PluginInitializeGorm{
								Type:        ast.TypePluginInitializeGorm,
								Path:        filepath.Join(global.RLB_CONFIG.AutoCode.Root, global.RLB_CONFIG.AutoCode.Server, "plugin", info.PackageName, secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext)),
								ImportPath:  fmt.Sprintf(`"%s/plugin/%s/model"`, global.RLB_CONFIG.AutoCode.Module, info.PackageName),
								StructName:  info.StructName,
								PackageName: "model",
								IsNew:       true,
							}
							asts[pluginInitializeGorm.Path+"=>"+pluginInitializeGorm.Type.String()] = pluginInitializeGorm
							creates[four] = pluginInitializeGorm.Path
						}
						if router != -1 {
							pluginInitializeRouter := &ast.PluginInitializeRouter{
								Type:                 ast.TypePluginInitializeRouter,
								Path:                 filepath.Join(global.RLB_CONFIG.AutoCode.Root, global.RLB_CONFIG.AutoCode.Server, "plugin", info.PackageName, secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext)),
								ImportPath:           fmt.Sprintf(`"%s/plugin/%s/router"`, global.RLB_CONFIG.AutoCode.Module, info.PackageName),
								AppName:              "Router",
								GroupName:            info.StructName,
								PackageName:          "router",
								FunctionName:         "Init",
								LeftRouterGroupName:  "public",
								RightRouterGroupName: "private",
							}
							asts[pluginInitializeRouter.Path+"=>"+pluginInitializeRouter.Type.String()] = pluginInitializeRouter
							creates[four] = pluginInitializeRouter.Path
						}
					}
				}
			}
		case "web":
		default:
			if templateDirs[i].Name() == ".DS_Store" {
				continue
			}
			return nil, nil, nil, errors.Errorf("[filpath:%s]非法模版文件!", second)
		}
	}

	return code, asts, creates, nil
}
