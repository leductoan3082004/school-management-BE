package handler

import (
	"SchoolManagement-BE/middleware"
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
}
