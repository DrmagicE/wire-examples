package main

import (
	"log"

	"github.com/DrmagicE/wire-examples/example/db"
	"github.com/DrmagicE/wire-examples/example/di"
	"github.com/DrmagicE/wire-examples/example/service"
)

// 手动注入依赖
/*func main() {
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
*/
// 使用wire自动注入依赖
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
