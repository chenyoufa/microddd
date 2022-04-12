package interfaces

import (
	"context"
	"microddd/application"

	"github.com/google/uuid"
)

type memberApi struct {
	*application.App
}

func NewMemberApi(app *application.App) *memberApi {
	return &memberApi{app}
}

func (mapi *memberApi) GetUser(uuid uuid.UUID) {
	mapi.Mapp.Get(context.Background(), uuid)
}
