package router

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	M api.memberApier
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
			gMenu.GET("", a.M.GetUser)

		}
	}
}
