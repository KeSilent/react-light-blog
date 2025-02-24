### 1. 安装 `gen` 工具

首先，确保你已经安装了 `gen` 工具。可以通过以下命令安装：

```bash
go install gorm.io/gen/tools/gentool@latest
```

安装完成后，可以通过 `gentool -h` 检查是否安装成功。

---

### 2. 配置 `gen` 工具

在项目根目录下创建一个配置文件（例如 `gen.yml` 或 `gen.go`），用于指定生成代码的参数。

#### 示例配置文件 (`gen.yml`)

```yaml
# gen.yml
db:
  dsn: 'user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local' # 数据库连接信息
  tables: # 指定需要生成的表，如果为空，则生成所有表
    - users
    - products
outPath: './dao' # 生成的代码输出目录
outFile: 'query.go' # 生成的查询代码文件名
modelPkgName: 'model' # 生成的模型代码包名
```
---

### 3. 运行 `gen` 生成代码

根据你的配置文件，运行 `gen` 工具生成代码。

#### 如果使用 `gen.yml` 配置文件：

```bash
gentool -c gen.yml
```
运行后，`gen` 会根据数据库表结构生成模型和查询代码，并保存到指定的目录中（例如 `./dao` 和 `./model`）。

---

### 4. 在项目中使用生成的代码

生成的代码通常包括：

- 模型代码（例如 `./model/user.go`）
- 查询代码（例如 `./dao/query.go`）

你可以在项目中直接使用这些生成的代码。例如：

```go
package main

import (
    "fmt"
    "your_project/dao"
    "your_project/model"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func main() {
    // 连接数据库
    dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // 初始化生成的查询对象
    q := dao.Use(db)

    // 使用生成的查询方法
    user, err := q.User.Where(q.User.ID.Eq(1)).First()
    if err != nil {
        panic("failed to query user")
    }

    fmt.Printf("User: %+v\n", user)
}
```

---

### 5. 结合 `AutoMigrate` 使用

如果你需要在项目启动时自动迁移数据库表结构，可以在 `main.go` 中结合 `AutoMigrate` 使用：

```go
package main

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "your_project/model" // 生成的模型包
)

func main() {
    // 连接数据库
    dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // 自动迁移数据库表结构
    err = db.AutoMigrate(&model.User{}, &model.Product{}) // 使用生成的模型
    if err != nil {
        panic("failed to migrate database")
    }

    // 其他业务逻辑
}
```

---

### 6. 自动化生成（可选）

如果你希望在每次数据库表结构变化时自动生成代码，可以将 `gen` 工具的运行命令添加到项目的构建脚本中，例如：

```bash
# build.sh
#!/bin/bash

# 生成代码
gentool -c gen.yml

# 构建项目
go build -o myapp
```

---

### 总结

通过以上步骤，你可以将 `gen` 工具集成到项目中，并结合 `AutoMigrate` 实现数据库表结构的自动迁移和代码生成。这种方式既能提高开发效率，又能确保代码与数据库结构的一致性。
