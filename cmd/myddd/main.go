package main

import (
	"fmt"
	api "myddd/interface/api"
	myroute "myddd/interface/router"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("init start ...")
	u := api.UserApi{}
	c := gin.Default()
	r := myroute.Router{}
	c.Routes(r.UserApi)
	c.Routes()

}
