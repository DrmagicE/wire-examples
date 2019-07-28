// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"
	"io"

	"github.com/DrmagicE/wire-examples/example/db"
	"github.com/DrmagicE/wire-examples/example/service"
	"github.com/google/wire"
)

// NewService 定义injector的函数签名
func NewService(c *db.Config, m *service.MailConfig) (*service.Service, error) {
	wire.Build(service.NewService, service.MailSet, service.UserSet, db.NewDB)
	return &service.Service{}, nil
}

func InitGreeter(ctx context.Context, s []string, w io.Writer) (*service.Greeter, error) {
	wire.Build(service.GreeterSet)
	return nil, nil
}
