//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
	"io"
	"os"
)

//injector parameter
func InitializedService(isError bool) (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository, NewSimpleService,
	)
	return nil, nil
}

//Multiple Binding
func InitializedDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDatabaseMongoDB, NewDatabasePostgreSQL, NewDatabaseRepository,
	)
	return nil
}

//provider set google wire for grouping
var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializedFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}

//binding interface google wire
var HelloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitializedHelloService() *HelloService {
	wire.Build(HelloSet, NewHelloService)
	return nil
}

//struct provider
func InitializedFooBar() *FooBar {
	wire.Build(
		NewFoo,
		NewBar,
		wire.Struct(new(FooBar), "Foo", "Bar"),
	)
	return nil
}

//binding values

var fooValue = &Foo{}
var barValue = &Bar{}

func InitializedFooBarValues() *FooBar {
	wire.Build(
		wire.Value(fooValue),
		wire.Value(barValue),
		wire.Struct(new(FooBar), "*"),
	)
	return nil
}

//injector interface value
func InitializedReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

//struct field provider
func InitializedConfiguration() *Configuration {
	wire.Build(
		NewApplication,
		wire.FieldsOf(new(*Application), "Configuration"),
	)
	return nil
}

//cleanup function
func InitializedConnection(name string) (*Connection, func()) {
	wire.Build(
		NewConnection,
		NewFile,
	)
	return nil, nil
}
