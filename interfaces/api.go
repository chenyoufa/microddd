package interfaces

import "microddd/application"

type Api struct {
	MApi memberApier
}

func NewApi(app *application.App) *Api {
	return &Api{
		MApi: &memberApi{app.Mapp},
	}
}
