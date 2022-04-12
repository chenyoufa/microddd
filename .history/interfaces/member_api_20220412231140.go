package interfaces

import "microddd/application"

type MemberApi struct {
	app *application.App
}

func NewMemberApi(app *application.ImgSearchApp) *SearchInf {
	return &SearchInf{app: app}
}
