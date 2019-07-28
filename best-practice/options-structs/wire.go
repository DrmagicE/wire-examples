// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"io"

	"context"

	"github.com/google/wire"
)

func InitializeGreeter(ctx context.Context, msg []Message, w io.Writer, r io.Reader) (*Greeter, error) {
	wire.Build(GreeterSet)
	return nil, nil
}
