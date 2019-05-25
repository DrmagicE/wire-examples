package service

// Service 处理业务
type Service struct {
	m MailSender
	u  UserRepository
}
// UserSignUp 用户注册，添加数据库，发邮件
func (s *Service) UserSignUp () {
	s.u.AddUser()
	s.m.Send()
}
// NewService provider：通过构造函数将MailSender和UserRepository注入
func NewService(m MailSender,u  UserRepository) *Service{
	return &Service{m:m, u:u}
}

