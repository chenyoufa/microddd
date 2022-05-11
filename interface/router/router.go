package router

import (

	// _ "microddd/interfaces/swagger"

	"myddd/interface/api"

	"github.com/gin-gonic/gin"
	// ginSwagger "github.com/swaggo/gin-swagger"
	// swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

// var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"))

type Router struct {
	*api.UserApi
}

func (a *Router) Register(app *gin.Engine) error {
	a.RegisterAPI(app)
	return nil
}
func (a *Router) RegisterAPI(app *gin.Engine) {
	// app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	g := app.Group("/api")
	// g.Use(middleware.UserAuthMiddleware(a.Auth))

	v1 := g.Group("/v1")
	{
		gMember := v1.Group("member")
		{
			gMember.GET("", a.Get)
			gMember.GET(":id", a.Get)
			gMember.POST("", a.Post)
			gMember.PUT("", a.Put)
			gMember.DELETE(":id", a.Del)
			// gMember.PATCH(":id/enable", a.MApi.Enable)
			// gMember.PATCH(":id/disable", a.MApi.Disable)
		}
	}
}
