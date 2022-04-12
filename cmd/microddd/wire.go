//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/google/wire"
	"microddd/infrastructure/db/dbcore"
	"microddd/infrastructure/repository"
)

//go:generate wire
var providerSet = wire.NewSet(
	// conf.NewViper,
	// conf.NewAppConfigCfg,
	// conf.NewLoggerCfg,
	// conf.NewRedisConfig,
	// conf.NewMongoConfig,
	// log.NewLogger,
	// redis.NewRedis,
	// mongo.NewMongo,

	dbcore.Connect,
	repository.NewRepository,
	// application.NewApps,
	// aggregate.NewFactory,
	// service.NewService,
	// adpter.NewSrv,
)

func NewApp() (*repository.Repository, error) {
	panic(wire.Build(providerSet))
}
