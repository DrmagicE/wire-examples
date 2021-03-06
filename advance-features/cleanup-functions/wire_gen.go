// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

// Injectors from wire.go:

func InitializeFileReader(path string) (*FileReader, func(), error) {
	fileReader, cleanup, err := NewFileReader(path)
	if err != nil {
		return nil, nil, err
	}
	return fileReader, func() {
		cleanup()
	}, nil
}
