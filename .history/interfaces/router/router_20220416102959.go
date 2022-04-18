package router

import (
	"microddd/interfaces/api"

	_ "microddd/interfaces/swagger"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"))

type Router struct {
	*api.WebApi
}

func (a *Router) Register(app *gin.Engine) error {
	a.RegisterAPI(app)
	return nil
}
func (a *Router) RegisterAPI(app *gin.Engine) {
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	g := app.Group("/api")
	// g.Use(middleware.UserAuthMiddleware(a.Auth))

	v1 := g.Group("/v1")
	{
		gMember := v1.Group("member")
		{
			gMember.GET("", a.MApi.QueryUser)
			gMember.GET("id", a.MApi.GetUser)
			gMember.POST("", a.MApi.CreateUser)
			gMember.PUT("/id", a.MApi.UpdateUser)
			gMember.DELETE("/id", a.MApi.DeleteUser)
			// gMember.PATCH(":id/enable", a.MApi.Enable)
			// gMember.PATCH(":id/disable", a.MApi.Disable)
		}
	}
}
