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

// @Tags GetUser
// @Summary 查询数据
// @Security ApiKeyAuth
// @Param current query int true "分页索引" default(1)
// @Param pageSize query int true "分页大小" default(10)
// @Param queryValue query string false "查询值"
// @Param status query int false "状态(1:启用 2:禁用)"
// @Success 200 {object} schema.ListResult{list=[]schema.Role} "查询结果"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/roles [get]
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
