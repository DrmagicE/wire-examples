// +build wireinject
// The build tag makes sure the stub is not built in the final build.


package di

import (
	"github.com/google/wire"
	"github.com/DrmagicE/wire-examples/service"
	"github.com/DrmagicE/wire-examples/db"
)

// NewService 定义injector的函数签名
func NewService(c *db.Config,m *service.MailConfig) (*service.Service,error) {
	wire.Build(service.NewService, service.MailSet, service.UserSet, db.NewDB)
	return &service.Service{},nil
}

