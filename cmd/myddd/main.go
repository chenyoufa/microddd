package main

import (
	"fmt"
	api "myddd/interface/api"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("init start ...")
	u := api.UserApi{}

	c := gin.Default()

	c.Routes()

}
