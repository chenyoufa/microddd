package api

import (
	"context"
	"microddd/application"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type memberApier interface {
	GetUser(c *gin.Context)
}

type memberApi struct {
	application.MemberApper
}

// var MemberApiSet = wire.NewSet(wire.Struct(new(memberApi), "*"))

func (mapi *memberApi) GetUser(c *gin.Context) {
	var err error
	uuid, err := uuid.Parse(c.Query("uuid"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "UnprocessableEntity")
	}
	dto, err := mapi.Get(context.Background(), uuid)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
	}
	c.JSON(http.StatusOK, dto)
}
