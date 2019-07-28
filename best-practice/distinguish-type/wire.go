// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import "github.com/google/wire"

// wire无法得知入参a,b跟FooBar.foo,FooBar.bar的对应关系
func InitializeFooBar(a Foo, b Bar) FooBar {
	wire.Build(NewFooBar)
	return FooBar{}
}
