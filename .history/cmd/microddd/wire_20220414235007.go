//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/google/wire"
	"microddd/application"
	"microddd/infrastructure/db/dbcore"
	"microddd/infrastructure/db/dbinit"
	"microddd/infrastructure/repository"
	"microddd/interfaces"
	// "microddd/interfaces/router"
)

//go:generate wire
var providerSet = wire.NewSet(
	dbinit.LoadConfig,
	dbcore.Connect,
	repository.NewRepository,
	application.NewApps,
	interfaces.NewApi,
	// router.RouterSet,
)

func NewAppw() (*interfaces.WebApi, error) {
	panic(wire.Build(providerSet))
}
