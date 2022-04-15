package api

import (
	"context"
	"microddd/application"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type memberApier interface {
	GetUser(uuid uuid.UUID)
}

type memberApi struct {
	application.MemberApper
}

// var MemberApiSet = wire.NewSet(wire.Struct(new(memberApi), "*"))

func (mapi *memberApi) GetUser(c gin.Context) {
	uuid := c.Query("uuid")
	mapi.Get(context.Background(), uuid)
}
