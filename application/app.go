package application

import (
	"microddd/domain/repository"
)

type App struct {
	Mapp MemberApper
}

func NewApps(repo *repository.AuthFactory) *App {
	return &App{
		&memberApp{mRepo: repo.MRepo},
	}
}

// var AppSet = wire.NewSet(
// 	memberAppSet,
// )

// var AppSet = wire.NewSet(
// 	wire.Struct(new(MyFoo)),
// 	wire.Bind(new(Fooer), new(MyFoo)))
