//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/google/wire"
	// "gorm.io/gorm"
	"microddd/application"
	"microddd/infrastructure/db/dbcore"
	"microddd/infrastructure/db/dbinit"
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
	dbinit.LoadConfig,
	dbcore.Connect,
	repository.NewRepository,
	application.NewApps,
	// service.NewService,
	// adpter.NewSrv,
)

func NewApp() (*application.App, error) {
	panic(wire.Build(providerSet))
}
