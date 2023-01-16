
# Facebook 开源 Golang 实体框架 Ent

![](https://cdn.jsdelivr.net/gh/ch3nnn/picgo/blog/imgs16738590832901.jpg)


ent是一个简单而又功能强大的Go语言实体框架，ent易于构建和维护应用程序与大数据模型。

- **图就是代码** - 将任何数据库表建模为Go对象。
- **轻松地遍历任何图形** - 可以轻松地运行查询、聚合和遍历任何图形结构。
- **静态类型和显式API** - 使用代码生成静态类型和显式API，查询数据更加便捷。
- **多存储驱动程序** - 支持MySQL, PostgreSQL, SQLite 和 Gremlin。
- **可扩展** - 简单地扩展和使用Go模板自定义。


## 安装
```console
go install entgo.io/ent/cmd/ent@latest
```
## 快速开始
```shell
go mod init <project>
```

### 创建你的第一个项目

进入你项目的根目录，然后运行：

```shell
go run -mod=mod entgo.io/ent/cmd/ent init User
```
以上的命令会在`<project>/ent/schema/`目录下产生`User`的数据模式（数据模式是数据库系统设计中的专业术语，若对该部分有任何理解问题，请查阅数据库系统的相关书籍）：

```go
package schema
import "entgo.io/ent"
// User在User实体中组合了ent默认的数据库模式定义
type User struct {
    ent.Schema
}
// User的字段
func (User) Fields() []ent.Field {
    return nil
}
// User的边
func (User) Edges() []ent.Edge {
    return nil
}
```

为`User` 模式添加两个字段：

```go
package schema
import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
)
// User的字段
func (User) Fields() []ent.Field {
    return []ent.Field{
        field.Int("age").
            Positive(),
        field.String("name").
            Default("unknown"),
    }
}
```

从项目的根目录下像如下命令那样，运行`go generate`：

```shell
go generate ./ent
```

上述命令，将产生如下的文件：
```text
ent
├── client.go
├── config.go
├── context.go
├── ent.go
├── generate.go
├── mutation.go
... truncated
├── schema
│   └── user.go
├── tx.go
├── user
│   ├── user.go
│   └── where.go
├── user.go
├── user_create.go
├── user_delete.go
├── user_query.go
└── user_update.go
```
### 创建你的第一个实体

首先，创建一个`ent.Client`。

> Sqlite3
```go
package main
import (
    "context"
    "log"
    "<project>/ent"
    _ "github.com/mattn/go-sqlite3"
)
func main() {
    client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
    if err != nil {
        log.Fatalf("failed opening connection to sqlite: %v", err)
    }
    defer client.Close()
    // 运行自动迁移工具。
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
}
```

> Postgres
```go
package main
import (
    "context"
    "log"
    "<project>/ent"
    _ "github.com/lib/pq"
)
func main() {
    client, err := ent.Open("postgres","host=<host> port=<port> user=<user> dbname=<database> password=<pass>")
    if err != nil {
        log.Fatalf("failed opening connection to postgres: %v", err)
    }
    defer client.Close()
    // 运行自动迁移工具。
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
}
```

> Mysql
```go
package main
import (
    "context"
    "log"
    "<project>/ent"
    _ "github.com/go-sql-driver/mysql"
)
func main() {
    client, err := ent.Open("mysql", "<user>:<pass>@tcp(<host>:<port>)/<database>?parseTime=True")
    if err != nil {
        log.Fatalf("failed opening connection to mysql: %v", err)
    }
    defer client.Close()
    // 运行自动迁移工具。
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
}
```
### 简单CURD
现在，我们准备创建我们的用户。 让我们写一个 `CreateUser` 函数，比如：
```go
// CreateUser 新建用户
func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.Create().
		SetAge(18).
		SetName("a8m").
		SetPhone("12345678").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil

}
```
`ent` 为每个实体结构生成一个package，包含其条件、默认值、验证器、有关存储元素的附加信息 (字段名、主键等) 。
```go
// QueryUser 查询用户
func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	// `Only` 如果没有发现用户则报错,
	// 否则正常返回。
	u, err := client.User.Query().Where(user.Name("a8m")).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

```

删除用户
```go
func DeleteUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, _ := QueryUser(ctx, client)
	if err := client.User.DeleteOne(u).Exec(ctx); err != nil {
		return nil, fmt.Errorf("failed delete user: %w", err)

	}
	return nil, nil
}
```
更新用户
```go
func UpdateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, _ := QueryUser(ctx, client)
	u, err := client.User.UpdateOne(u).SetAge(28).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed update user: %w", err)
	}
	return u, nil

}
```

## 参考
[Ent 中文文档](https://ent.ryansu.pro/#/zh-cn/getting-started)

[Ent Github](https://github.com/ent/ent)