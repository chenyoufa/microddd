package main

import "github.com/google/uuid"

func main() {
	api, err := NewApp()
	if err != nil {
		panic(err)
	}
	api.MApi.GetUser(uuid.New())

	// var abp application.MemberApper
	// config, _ := dbinit.LoadConfig()
	// db, _ := dbcore.Connect(config)
	// factory := repository.NewRepository(db)
	// abp = application.NewApps(factory).Mapp

	// abp.Get(context.Background(), uuid.New())

}
