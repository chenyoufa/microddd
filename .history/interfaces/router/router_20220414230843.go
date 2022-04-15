package router

import (
	"microddd/interfaces/api"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"))

type Router struct {
	api.WebApi
}

func (a *Router) Register(app *gin.Engine) error {
	a.RegisterAPI(app)
	return nil
}
func (a *Router) RegisterAPI(app *gin.Engine) {
	g := app.Group("/api")
	// g.Use(middleware.UserAuthMiddleware(a.Auth))

	v1 := g.Group("/v1")
	{
		gMenu := v1.Group("menus")
		{
			gMenu.GET("", a.MApi.GetUser)

		}
	}
}
