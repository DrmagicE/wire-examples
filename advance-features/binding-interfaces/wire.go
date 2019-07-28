// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import "github.com/google/wire"

func InitializeUserService(foo string, bar int) (*UserService){
	wire.Build(NewUserService,MockUserRepoSet)
	return nil
}
