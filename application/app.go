package application

import "microddd/domain/repository"

type App struct {
	mapp *memberApp
}

func NewApps(repo repository.MemberRepo) *App {
	return &App{
		&memberApp{mRepo: repo},
	}
}
