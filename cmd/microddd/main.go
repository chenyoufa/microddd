package main

import (
	"context"
	"microddd/application"
	"microddd/infrastructure/db/dbcore"
	"microddd/infrastructure/db/dbinit"
	"microddd/infrastructure/repository"

	"github.com/google/uuid"
)

func main() {
	// app, err := NewApp()
	// if err != nil {
	// 	panic(err)
	// }
	// app.Get()

	config, _ := dbinit.LoadConfig()
	db, _ := dbcore.Connect(config)
	factory := repository.NewRepository(db)
	app := application.NewApps(factory)
	app.Mapp.Get(context.Background(), uuid.New())

}
