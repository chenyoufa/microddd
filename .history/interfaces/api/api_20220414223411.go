package api

import "microddd/application"

type WebApi struct {
	MApi *memberApier
}

func NewApi(app *application.App) *WebApi {
	return &WebApi{
		MApi: &memberApi{app.Mapp},
	}
}
