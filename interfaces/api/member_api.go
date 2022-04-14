package api

import (
	"context"
	"microddd/application"

	"github.com/google/uuid"
)

type memberApier interface {
	GetUser(uuid uuid.UUID)
}

type memberApi struct {
	application.MemberApper
}

// var MemberApiSet = wire.NewSet(wire.Struct(new(memberApi), "*"))

func (mapi *memberApi) GetUser(uuid uuid.UUID) {
	mapi.Get(context.Background(), uuid)
}
