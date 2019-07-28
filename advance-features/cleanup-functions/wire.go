// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import "github.com/google/wire"

func InitializeFileReader(path string) (*FileReader,func(), error) {
	wire.Build(NewFileReader)
	return nil,nil,nil
}
