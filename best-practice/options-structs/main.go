package main

import (
	"io"

	"context"

	"github.com/google/wire"
)

type Message string

// Options
type Options struct {
	Messages []Message
	Writer   io.Writer
	Reader   io.Reader
}
type Greeter struct {
}

// NewGreeter 使用Options以避免构造函数过长
func NewGreeter(ctx context.Context, opts *Options) (*Greeter, error) {
	return nil, nil
}

// GreeterSet 使用wire.Struct设置Options为provider
var GreeterSet = wire.NewSet(wire.Struct(new(Options), "*"), NewGreeter)
