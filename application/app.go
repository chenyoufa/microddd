package application

import "microddd/domain/repository"

type App struct {
	Mapp *memberApp
}

func NewApps(repo *repository.AuthFactory) *App {
	return &App{
		&memberApp{mRepo: repo.MRepo},
	}
}
