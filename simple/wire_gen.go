// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package simple

// Injectors from injector.go:

func InitializeService() *SimpleService {
	simpleRepository := NewSimpleRepository()
	simpleService := NewSimpleService(simpleRepository)
	return simpleService
}
