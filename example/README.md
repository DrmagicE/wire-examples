# example

[wire](https://github.com/google/wire)进行依赖注入使用示例。本示例以用户注册场景为例，用户注册成功之后：
- 保存到数据库 
- 下发注册成功邮件 

抽象两个接口：
- `service.UserRepository`表示数据库依赖
- `service.MailSender`表示邮件依赖  


# Prerequisite
先安装wire：https://github.com/google/wire#installing
# Installation
```
go get github.com/DrmagicE/wire-examples
```

# wire provider
```
// service/email.go 

// NewMailSender provider: 根据邮件配置初始化 mailSender
func NewMailSender(m *MailConfig) *mailSender {
	return &mailSender{}
}

// MailSet 声明NewMailSender的返回值是MailSender接口类型
var MailSet = wire.NewSet(NewMailSender, wire.Bind(new(MailSender), new(*mailSender)))
```

```
// service/user.go 

// NewUserRepo provider: 根据*sql.DB初始化 *userRepo
func NewUserRepo(db *sql.DB) *userRepo{
	return &userRepo{}
}

// UserSet 声明NewUserRepo的返回值是UserRepository接口类型
var UserSet = wire.NewSet(NewUserRepo, wire.Bind(new(UserRepository), new(*userRepo)))
```

```
// service/service.go 

// NewService provider：通过构造函数将MailSender和UserRepository注入
func NewService(m MailSender,u  UserRepository) *Service{
	return &Service{m:m, u:u}
}
```

由于`NewService`接收的是接口类型，在wire中，需要通过`wire.NewSet`来将接口与结构体绑定。详见：https://github.com/google/wire/blob/master/docs/guide.md#binding-interfaces

# 定义wire injector
```
// di/wire.go

// NewService 定义injector的函数签名
func NewService(c *db.Config,m *service.MailConfig) (*service.Service,error) {
	wire.Build(service.NewService, service.MailSet, service.UserSet, db.NewDB)
	return &service.Service{},nil
}
```


# 生成injector
```
$ cd di
$ wire
```

# 运行
```
go run main.go
```

# 手动注入 VS wire注入
手动注入：
```
// 手动注入依赖
func main() {
    // 数据库配置
	dbConfig := &db.Config{}
	// 邮件配置
	mailConfig := &service.MailConfig{}
	// sql 依赖数据库配置
	sql, err := db.NewDB(dbConfig)
	if err != nil {
		log.Fatal(err)
	}
	// userRepo 依赖 sql
	userRepo := service.NewUserRepo(sql)
	// mailSender 依赖 mailConfig
	mailSender := service.NewMailSender(mailConfig)
	s := service.NewService(mailSender,userRepo)
	s.UserSignUp()
}
```
wire注入：
```
// 使用wire生成的代码注入依赖
func main() {
	// 邮件配置
	dbConfig := &db.Config{}
	// 邮件配置
	mailConfig := &service.MailConfig{}
	s, err := di.NewService(dbConfig, mailConfig)
	if err != nil {
		log.Fatal(err)
	}
	s.UserSignUp()
}
```
使用`wire`之后，`main`变得清爽多了，依赖越复杂，`wire`的效果越明显。
