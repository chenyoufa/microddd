package api

import (
	"context"
	"fmt"
	"log"
	"microddd/application"
	"microddd/application/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type memberApier interface {
	GetUser(c *gin.Context)
	QueryUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	// Enable(c *gin.Context)
	// Disable(c *gin.Context)
}

type memberApi struct {
	application.MemberApper
}

// @Tags UserAPI
// @Summary 查询指定数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Success 200
// @Router /api/v1/member/{id} [get]
func (mapi *memberApi) GetUser(c *gin.Context) {
	var err error
	uuid, err := strconv.ParseUint(c.Param("id"), 10, 64)
	log.Println("uuid:", uuid)
	c.JSON(http.StatusOK, uuid)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "UnprocessableEntity")
	}
	dto, err := mapi.Get(context.Background(), uuid)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
	}
	c.JSON(http.StatusOK, dto)
}

// @Tags UserAPI
// @Summary 查询指定数据列表
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Success 200
// @Router /api/v1/member [get]
func (mapi *memberApi) QueryUser(c *gin.Context) {
	var err error
	var mdto dto.Member_dto
	if err := c.ShouldBindJSON(&mdto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Login information is not complete"})
		return
	}
	dto, err := mapi.GetList(context.Background(), mdto.ID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
	}
	c.JSON(http.StatusOK, dto)
}

// @Tags UserAPI
// @Summary 新增数据
// @Security ApiKeyAuth
// @Param  body body dto.Member_dto true "交款查询参数"
// @Success 200 {object} bool
// @Router /api/v1/member [post]
func (mapi *memberApi) CreateUser(c *gin.Context) {
	fmt.Println("print:")
	var err error
	var mdto dto.Member_dto

	if err := c.ShouldBindJSON(&mdto); err != nil {
		fmt.Println("err1:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Login information is not complete"})
		return
	}
	fmt.Println("mdto:", mdto)
	dto, err := mapi.Add(context.Background(), &mdto)
	if err != nil {
		fmt.Println("err2:", err)
		c.JSON(http.StatusUnauthorized, err)
	}
	c.JSON(http.StatusOK, dto)
}

// @Tags UserAPI
// @Summary 更新数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Success 200
// @Router /api/v1/member/{id} [put]
func (mapi *memberApi) UpdateUser(c *gin.Context) {
	var err error
	var mdto dto.Member_dto
	if err := c.ShouldBindJSON(&mdto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Login information is not complete"})
		return
	}
	dto, err := mapi.Edit(context.Background(), &mdto)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
	}
	c.JSON(http.StatusOK, dto)
}

// @Tags UserAPI
// @Summary 删除指定数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Success 200
// @Router /api/v1/member/{id} [delete]
func (mapi *memberApi) DeleteUser(c *gin.Context) {
	var err error
	uuid := c.Query("uuid")
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "UnprocessableEntity")
	}
	dto, err := mapi.Remove(context.Background(), uuid)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
	}
	c.JSON(http.StatusOK, dto)
}
