package core

import (
	"os"

	"github.com/kesilent/react-light-blog/global"
	"github.com/kesilent/react-light-blog/utils"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

// GenStructs 生成dal/query及model目录下的结构体
func GenStructs() error {
	dir, _ := os.Getwd()
	// 指定生成代码的具体相对目录(相对当前文件)，默认为：./query
	// 默认生成需要使用WithContext之后才可以查询的代码，但可以通过设置gen.WithoutContext禁用该模式
	generator := gen.NewGenerator(gen.Config{
		// 默认会在 OutPath 目录生成CRUD代码，并且同目录下生成 model 包
		// 所以OutPath最终package不能设置为model，在有数据库表同步的情况下会产生冲突
		// 若一定要使用可以通过ModelPkgPath单独指定model package的名称
		OutPath: dir + "/dal/query",
		/* ModelPkgPath: "dal/model"*/

		// gen.WithoutContext：禁用WithContext模式
		// gen.WithDefaultQuery：生成一个全局Query对象Q
		// gen.WithQueryInterface：生成Query接口
		Mode: gen.WithDefaultQuery,

		// 表字段可为 null 值时, 对应结体字段使用指针类型
		FieldNullable: false, // generate pointer when field is nullable

		// 表字段默认值与模型结构体字段零值不一致的字段, 在插入数据时需要赋值该字段值为零值的, 结构体字段须是指针类型才能成功, 即`FieldCoverable:true`配置下生成的结构体字段.
		// 因为在插入时遇到字段为零值的会被GORM赋予默认值. 如字段`age`表默认值为10, 即使你显式设置为0最后也会被GORM设为10提交.
		// 如果该字段没有上面提到的插入时赋零值的特殊需要, 则字段为非指针类型使用起来会比较方便.
		FieldCoverable: false, // generate pointer when field has default value, to fix problem zero value cannot be assign: https://gorm.io/docs/create.html#Default-Values

		// 模型结构体字段的数字类型的符号表示是否与表字段的一致, `false`指示都用有符号类型
		FieldSignable: false, // detect integer field's unsigned type, adjust generated data type
		// 生成 gorm 标签的字段索引属性
		FieldWithIndexTag: false, // generate with gorm index tag
		// 生成 gorm 标签的字段类型属性
		FieldWithTypeTag: true, // generate with gorm column type tag
	})
	// 通常复用项目中已有的SQL连接配置db(*gorm.DB)
	// 非必需，但如果需要复用连接时的gorm.Config或需要连接数据库同步表信息则必须设置
	generator.UseDB(global.RLB_DB)

	generator.WithJSONTagNameStrategy(func(columnName string) (tagContent string) {
		return utils.LowerCamelCase(columnName)
	})

	// 自定义字段的数据类型
	// 统一数字类型为int64,兼容protobuf和thrift
	dataMap := map[string]func(detailType gorm.ColumnType) (dataType string){
		"varchar": func(detailType gorm.ColumnType) (dataType string) { return "string" },
		"tinyint": func(detailType gorm.ColumnType) (dataType string) { return "bool" },
		"timestamp": func(detailType gorm.ColumnType) (dataType string) {
			// 只将 update_time 和 deleted_at 设置为指针类型
			columnName := detailType.Name()
			if columnName == "update_time" || columnName == "deleted_at" {
				return "*time.Time"
			}
			return "time.Time"
		},
		"datetime": func(detailType gorm.ColumnType) (dataType string) {
			// 只将 update_time 和 deleted_at 设置为指针类型
			columnName := detailType.Name()
			if columnName == "update_time" || columnName == "deleted_at" {
				return "*time.Time"
			}
			return "time.Time"
		},
	}
	// 要先于`ApplyBasic`执行
	generator.WithDataTypeMap(dataMap)

	fieldOpts := []gen.ModelOpt{}
	// 将非默认字段名的字段定义为自动时间戳和软删除字段;
	// 自动时间戳默认字段名为:`updated_at`、`created_at, 表字段数据类型为: INT 或 DATETIME
	// 软删除默认字段名为:`deleted_at`, 表字段数据类型为: DATETIME

	allModel := generator.GenerateAllTable()
	generator.ApplyBasic(allModel...)
	// 1. 首先生成基础模型
	SysAuthorityMenus := generator.GenerateModel("sys_authority_menus")
	SysUserAuthority := generator.GenerateModel("sys_user_authority")

	// 生成角色模型
	authority := generator.GenerateModel("sys_authorities")

	// 生成菜单模型
	menu := generator.GenerateModel("sys_base_menus")

	// 重新配置用户模型的关联关系
	user := generator.GenerateModel("sys_users", append(fieldOpts,
		gen.FieldRelate(field.Many2Many, "Authorities", authority, &field.RelateConfig{
			GORMTag: field.GormTag{
				"many2many":      []string{"sys_user_authority"},
				"joinForeignKey": []string{"sys_user_id"},
				"joinReferences": []string{"sys_authority_authority_id"},
			},
		}),
		gen.FieldRelate(field.BelongsTo, "Authority", authority, &field.RelateConfig{
			GORMTag: field.GormTag{"foreignKey": []string{"authority_id"}},
		}),
	)...)

	// 重新配置角色模型的关联关系
	authority = generator.GenerateModel("sys_authorities", append(fieldOpts,
		gen.FieldRelate(field.Many2Many, "Menus", menu, &field.RelateConfig{
			GORMTag: field.GormTag{
				"many2many":      []string{"sys_authority_menus"},
				"joinForeignKey": []string{"sys_authority_authority_id"},
				"joinReferences": []string{"sys_base_menu_id"},
			},
		}),
		gen.FieldRelate(field.Many2Many, "Users", user, &field.RelateConfig{
			GORMTag: field.GormTag{
				"many2many":      []string{"sys_user_authority"},
				"joinForeignKey": []string{"sys_authority_authority_id"},
				"joinReferences": []string{"sys_user_id"},
			},
		}),
	)...)

	// 重新配置菜单模型的关联关系
	menu = generator.GenerateModel("sys_base_menus", append(fieldOpts,
		gen.FieldRelate(field.Many2Many, "Authorities", authority, &field.RelateConfig{
			GORMTag: field.GormTag{
				"many2many":      []string{"sys_authority_menus"},
				"joinForeignKey": []string{"sys_base_menu_id"},
				"joinReferences": []string{"sys_authority_authority_id"},
			},
		}),
		gen.FieldRelate(field.HasMany, "Children", menu, &field.RelateConfig{
			GORMTag: field.GormTag{"foreignKey": []string{"parent_id"}},
		}),
	)...)

	// 生成字典相关模型
	dicDetails := generator.GenerateModel("sys_dictionary_details")

	// 生成带有关联关系的字典模型
	dic := generator.GenerateModel("sys_dictionaries", append(fieldOpts,
		gen.FieldRelate(field.HasMany, "Details", dicDetails, &field.RelateConfig{
			GORMTag: field.GormTag{"foreignKey": []string{"sys_dictionary_id"}},
		}),
	)...)

	// 重新生成带有关联关系的字典详情模型
	dicDetails = generator.GenerateModel("sys_dictionary_details", append(fieldOpts,
		gen.FieldRelate(field.BelongsTo, "Dictionary", dic, &field.RelateConfig{
			GORMTag: field.GormTag{
				"foreignKey": []string{"sys_dictionary_id"},
				"references": []string{"id"},
			},
		}),
	)...)

	// 应用所有模型
	generator.ApplyBasic(
		SysAuthorityMenus,
		SysUserAuthority,
		user,
		authority,
		menu,
		dic,
		dicDetails,
	)

	// 执行并生成代码
	generator.Execute()
	return nil
}
