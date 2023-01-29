// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package simple

import (
	"github.com/google/wire"
)

// Injectors from injector.go:

func InitializeService(isError bool) (*SimpleService, error) {
	simpleRepository := NewSimpleRepository(isError)
	simpleService, err := NewSimpleService(simpleRepository)
	if err != nil {
		return nil, err
	}
	return simpleService, nil
}

func InitializeDatabaseRepository() *DatabaseRepository {
	databasePostgresql := NewDatabasePostgresql()
	databaseMongodb := NewDatabaseMongodb()
	databaseRepository := NewDatabaseRepository(databasePostgresql, databaseMongodb)
	return databaseRepository
}

func InitializeFoobarService() *FoobarService {
	fooRepository := NewFooRepository()
	fooService := NewFooService(fooRepository)
	barRepository := NewBarRepository()
	barService := NewBarService(barRepository)
	foobarService := NewFoobarService(fooService, barService)
	return foobarService
}

// injector.go:

var fooset = wire.NewSet(NewFooRepository, NewFooService)

var barset = wire.NewSet(NewBarRepository, NewBarService)
