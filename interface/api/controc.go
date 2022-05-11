package api

import (
	"fmt"
	"myddd/application/user"
	"myddd/domain/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
	App user.UserApp
}

func (uapi *UserApi) Get(c *gin.Context) {
	id := c.Query("id")
	rtval := uapi.App.Query(id)
	c.JSON(http.StatusOK, rtval)
}
func (uapi *UserApi) Post(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")
	nickname := c.Query("nickname")
	model, err := model.NewUser(name, password, nickname)
	if err != nil {
		fmt.Println("err:%v", err)
		c.JSON(http.StatusInternalServerError, err)
	}
	rtval := uapi.App.Add(model)
	c.JSON(http.StatusOK, rtval)

}
func (uapi *UserApi) Del(c *gin.Context) {
	id := c.Query("id")
	rtval := uapi.App.Del(id)
	c.JSON(http.StatusOK, rtval)
}

//疑问
func (uapi *UserApi) Put(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 32, 64)
	name := c.Query("name")
	password := c.Query("password")
	nickname := c.Query("nickname")
	status, err := strconv.ParseInt(c.Query("status"), 32, 64)

	userdo := &model.User{id, name, password, nickname, status}
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
	}
	rtval := uapi.App.Update(userdo)
	c.JSON(http.StatusOK, rtval)
}
