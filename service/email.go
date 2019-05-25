package service

import (
	"log"

	"github.com/google/wire"
)

// MailSender 发送邮件接口
type MailSender interface {
	Send()
}


// MailConfig 邮件配置
type MailConfig struct {
}
// mailSender MailSender接口实现
type mailSender struct {

}

// Send 发邮件
func (e *mailSender) Send() {
	log.Println("send email")
}

// NewMailSender provider: 根据邮件配置初始化 mailSender
func NewMailSender(m *MailConfig) *mailSender {
	return &mailSender{}
}

// MailSet 声明NewMailSender的返回值是MailSender接口类型
var MailSet = wire.NewSet(NewMailSender, wire.Bind(new(MailSender), new(*mailSender)))