package service

// Service 处理业务
type Service struct {
	m MailSender
	u  UserRepository

	handler h
}

type Handler struct {

}
func (H *Handler) DoA() {

}
func (H *Handler) DoB() {

}

type h interface {
	DoA()
	DoB()
}
// UserSignUp 用户注册，添加数据库，发邮件
func (s *Service) UserSignUp () {
	s.u.AddUser()
	s.m.Send()
	s.handler.DoA()
}


// NewService provider：通过构造函数将MailSender和UserRepository注入
func NewService(m MailSender,u  UserRepository) *Service{
	return &Service{m:m, u:u, handler:&Handler{}}
}

