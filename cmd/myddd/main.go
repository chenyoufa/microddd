package main

import (
	"fmt"
	app "myddd/application/user"
	repo "myddd/infrastructure/repositores"
	api "myddd/interface/api"
	route "myddd/interface/router"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("init start ...")
	myrepo := repo.NewUserPo("127.0.0.1:3306")

	// var myrepo = &repo.UserPo{db}
	myapp := app.UserApp{Repo: myrepo}
	myapi := &api.UserApi{App: myapp}

	engine := gin.Default()
	r := route.Router{
		UserApi: myapi,
	}
	r.Register(engine)

	engine.Run()
}
