//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func InitializeService(isError bool) (*SimpleService, error) {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil, nil
}

func InitializeDatabaseRepository() *DatabaseRepository {
	wire.Build(NewDatabasePostgresql, NewDatabaseMongodb, NewDatabaseRepository)
	return nil
}

var fooset = wire.NewSet(NewFooRepository, NewFooService)
var barset = wire.NewSet(NewBarRepository, NewBarService)

func InitializeFoobarService() *FoobarService {
	wire.Build(fooset, barset, NewFoobarService)
	return nil
}

var helloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitializeHelloService() *HelloService {
	wire.Build(helloSet, NewHelloService)
	return nil

}

var foobarSet = wire.NewSet(
	NewFoo,
	NewBar,
)

func InitializeFoobar() *FooBar {
	wire.Build(foobarSet, wire.Struct(new(FooBar), "Foo", "Bar")) // * for all fields
	return nil
}
