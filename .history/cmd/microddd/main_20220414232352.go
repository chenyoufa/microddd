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
	// _, err := NewApp()
	// if err != nil {
	// 	panic(err)
	// }

	var abp application.MemberApper
	config, _ := dbinit.LoadConfig()
	db, _ := dbcore.Connect(config)
	factory := repository.NewRepository(db)
	abp = application.NewApps(factory).Mapp

	abp.Get(context.Background(), uuid.New())

	// model := aggregate.Member_aggre{}

	// var ids []uuid.UUID
	// u1, _ := uuid.Parse("17dd5571-ecee-4e01-99a3-86e1cdd46163")
	// u2, _ := uuid.Parse("4e753585-4f55-45a1-9f96-ab6b97631e42")
	// u3, _ := uuid.Parse("28eee28f-3732-4193-97cf-55b1d7ecdca3")
	// ids = append(ids, u1)
	// ids = append(ids, u2)
	// ids = append(ids, u3)

	// user := entity.UserEntity{
	// 	uuid.New(),
	// 	"",
	// 	"",
	// 	"",
	// }
	// model.User = &user
	// model.AddRoles(ids...)

	// model.RemoveRoles(u1)

	// model.Delete()
	// fmt.Println(model.GetRoleIDs())
	// app := gin.Default()
	// router := router.Router{api}
	// router.Register(app)
	// app.Run()

}
