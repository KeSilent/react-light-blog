package system

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/kesilent/react-light-blog/dal/request"
	response "github.com/kesilent/react-light-blog/dal/response"
	"github.com/kesilent/react-light-blog/global"
	"github.com/kesilent/react-light-blog/services/autocode"
	utilsAst "github.com/kesilent/react-light-blog/utils/ast"
	utilsAutocode "github.com/kesilent/react-light-blog/utils/autocode"
	"github.com/pkg/errors"
)

var AutoCodeServiceApp = new(AutoCodeService)

type AutoCodeService struct {
}

/**
 * @Author: Yang
 * @description: 获取数据库中所有的表
 * @return {*}
 */
func (autoCodeService *AutoCodeService) GetAllTableName() (tables []string, err error) {

	// 获取当前数据库名称
	var dbName string
	global.RLB_DB.Raw("SELECT DATABASE()").Scan(&dbName)

	// 查询所有表
	err = global.RLB_DB.Raw("SHOW TABLES FROM `" + dbName + "`").Scan(&tables).Error
	if err != nil {
		return nil, err
	}

	return tables, nil
}

/**
 * @Author: Yang
 * @description: 获取表中所有的字段
 * @param {string} tableName
 * @return {*}
 */
func (autoCodeService *AutoCodeService) GetFieldsByTableName(tableName string) (fields []response.AutoCodeField, err error) {

	// 查询所有表
	err = global.RLB_DB.Raw("SHOW FULL COLUMNS FROM `" + tableName + "`").Scan(&fields).Error
	if err != nil {
		return nil, err
	}

	return fields, nil
}

/**
 * @Author: Yang
 * @description: 创建生成自动化代码
 * @param {request.AutoCodeReq} autoCodeReq
 * @return {*}
 */
func (autoCodeService *AutoCodeService) CreateTemp(autoCodeReq request.AutoCodeReq) (err error) {
	//如果判断需要自动创建资源标示：AutoCreateResource，则先为数据表增加相应字段CreatedBy等

	//生成模拟数据

	autoCodeReq = request.AutoCodeReq{
		PackageName:        "demo",
		Module:             global.RLB_CONFIG.AutoCode.Module,
		TableName:          "user",
		AutoCreateResource: true,
		Relations: []request.RelationReq{
			{
				RelateTable:   "role",
				RelateType:    "many2one",
				RelateColumn:  "role_id",
				RieldName:     "Role",
				RelationTable: "user_role",
			},
		},
		Template:        "plugin",
		GenerateServer:  true,
		StructName:      "User",
		Abbreviation:    "user",
		HumpPackageName: "User",
		Description:     "用户",
	}

	generate, _, _, err := autoCodeService.generate(autoCodeReq)
	if err != nil {
		return err
	}
	for key, builder := range generate {
		err = os.MkdirAll(filepath.Dir(key), os.ModePerm)
		if err != nil {
			return errors.Wrapf(err, "[filepath:%s]创建文件夹失败!", key)
		}
		err = os.WriteFile(key, []byte(builder.String()), 0666)
		if err != nil {
			return errors.Wrapf(err, "[filepath:%s]写入文件失败!", key)
		}
	}

	return nil
}

func (autoCodeService *AutoCodeService) generate(info request.AutoCodeReq) (map[string]strings.Builder, map[string]string, map[string]utilsAst.Ast, error) {

	templates, asts, _, err := autocode.AutoCodePackage.Templates(info, false)
	if err != nil {
		return nil, nil, nil, err
	}
	code := make(map[string]strings.Builder)
	for key, create := range templates {
		var files *template.Template
		files, err = template.New(filepath.Base(key)).Funcs(utilsAutocode.GetTemplateFuncMap()).ParseFiles(key)
		if err != nil {
			return nil, nil, nil, errors.Wrapf(err, "[filpath:%s]读取模版文件失败!", key)
		}
		var builder strings.Builder
		err = files.Execute(&builder, info)
		if err != nil {
			return nil, nil, nil, errors.Wrapf(err, "[filpath:%s]生成文件失败!", create)
		}
		code[create] = builder
	} // 生成文件

	injections := make(map[string]utilsAst.Ast, len(asts))
	for key, value := range asts {
		keys := strings.Split(key, "=>")
		if len(keys) == 2 {
			if keys[1] == utilsAst.TypePluginInitializeV2 {
				continue
			}
			if keys[1] == utilsAst.TypePackageInitializeGorm || keys[1] == utilsAst.TypePluginInitializeGorm {
				continue
			}
			var builder strings.Builder
			parse, _ := value.Parse("", &builder)
			if parse != nil {
				_ = value.Injection(parse)
				err = value.Format("", &builder, parse)
				if err != nil {
					return nil, nil, nil, err
				}
				code[keys[0]] = builder
				injections[keys[1]] = value
				fmt.Println(keys[0], "注入成功!")
			}
		}
	}
	// 注入代码
	return code, templates, injections, nil
}
