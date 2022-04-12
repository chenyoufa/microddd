package interfaces

import "microddd/application"

type memberApi struct {
	app *application.App
}

func NewMemberApi(app *application.App) *memberApi {
	return &memberApi{app: app}
}
