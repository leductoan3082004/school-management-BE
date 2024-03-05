package handler

import (
	"SchoolManagement-BE/middleware"
	coursegin "SchoolManagement-BE/modules/course/transport/gin"
	usergin "SchoolManagement-BE/modules/user/transport/gin"
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
)

func MainRoute(router *gin.Engine, sc goservice.ServiceContext) {
	router.Use(middleware.AllowCORS())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Use(middleware.Recover())

	v1 := router.Group("/v1")
	v1.POST("user/login", usergin.Login(sc))

	authedRoutes := v1.Group("/", middleware.RequiredAuth(sc))
	authedRoutes.GET("/user", usergin.GetProfile(sc))
	authedRoutes.POST("/user", middleware.AdminAuthorization(), usergin.Create(sc))
	authedRoutes.POST("/user/change-password", usergin.ChangePassword(sc))

	course := v1.Group("/course")
	{
		course.POST(
			"/",
			middleware.RequiredAuth(sc),
			middleware.AdminAuthorization(),
			coursegin.Create(sc),
		)
		course.DELETE(
			"/",
			middleware.RequiredAuth(sc),
			middleware.AdminAuthorization(),
			coursegin.Delete(sc),
		)
		course.PUT(
			"/",
			middleware.RequiredAuth(sc),
			middleware.AdminAuthorization(),
			coursegin.Update(sc),
		)

		course.GET("/", coursegin.List(sc))
	}
}
