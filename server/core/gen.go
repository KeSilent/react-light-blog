package core

import (
	"fmt"
	"os"

	"github.com/kesilent/react-light-blog/global"
	"github.com/kesilent/react-light-blog/utils"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func getGenerator(packageName string) *gen.Generator {
	if packageName == "" {
		return nil
	}

	dir, _ := os.Getwd()
	outPath := dir + "/dal/query"
	if _, err := os.Stat(outPath); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			fmt.Println("创建目录失败：", err)
		} else {
			fmt.Println("目录创建成功")
		}
	}

	// 指定生成代码的具体相对目录(相对当前文件)，默认为：./query
	// 默认生成需要使用WithContext之后才可以查询的代码，但可以通过设置gen.WithoutContext禁用该模式
	generator := gen.NewGenerator(gen.Config{
		// 默认会在 OutPath 目录生成CRUD代码，并且同目录下生成 model 包
		// 所以OutPath最终package不能设置为model，在有数据库表同步的情况下会产生冲突
		// 若一定要使用可以通过ModelPkgPath单独指定model package的名称
		OutPath: outPath,
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
		"tinyint": func(detailType gorm.ColumnType) (dataType string) { return "*bool" },
		// "smallint":  func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		// "mediumint": func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		// "bigint":    func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		// "int":       func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		// "timestamp": func(detailType gorm.ColumnType) (dataType string) { return "LocalTime" }, // 自定义时间
		// "decimal":   func(detailType gorm.ColumnType) (dataType string) { return "Decimal" },   // 金额类型全部转换为第三方库,github.com/shopspring/decimal
	}
	// 要先于`ApplyBasic`执行
	generator.WithDataTypeMap(dataMap)
	return generator
}
