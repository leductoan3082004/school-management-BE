package handler

import (
	"SchoolManagement-BE/middleware"
	classroomgin "SchoolManagement-BE/modules/classroom/transport/gin"
	coursegin "SchoolManagement-BE/modules/course/transport/gin"
	lessongin "SchoolManagement-BE/modules/lesson/transport/gin"
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

	user := authedRoutes.Group("/user")
	{
		user.GET("/profile", usergin.GetProfile(sc))
		user.POST("/", middleware.AdminAuthorization(), usergin.Create(sc))
		user.POST("/change-password", usergin.ChangePassword(sc))
		user.GET("/", middleware.AdminAuthorization(), usergin.ListUsers(sc))
		user.GET("/class", usergin.ListJoinedClass(sc))
	}

	course := authedRoutes.Group("/course")
	{
		course.POST(
			"/",
			middleware.AdminAuthorization(),
			coursegin.Create(sc),
		)
		course.DELETE(
			"/",
			middleware.AdminAuthorization(),
			coursegin.Delete(sc),
		)
		course.PUT(
			"/",
			middleware.AdminAuthorization(),
			coursegin.Update(sc),
		)

		course.GET("/", coursegin.List(sc))
	}

	classroom := authedRoutes.Group("/classroom")
	{
		classroom.POST(
			"/",
			middleware.AdminAuthorization(),
			classroomgin.Create(sc),
		)
		classroom.DELETE(
			"/",
			middleware.AdminAuthorization(),
			classroomgin.Delete(sc),
		)
		classroom.GET("/", classroomgin.List(sc))
		classroom.PUT(
			"/",
			middleware.AdminAuthorization(),
			classroomgin.Update(sc),
		)

		member := classroom.Group("/member")
		{
			member.POST(
				"/",
				middleware.AdminAuthorization(),
				classroomgin.AddMemberToClass(sc),
			)
			member.POST("/:class_id", classroomgin.StudentRegisterClass(sc))
			member.GET(
				"/",
				middleware.TeacherValidationForm(sc),
				classroomgin.ListMemberInClass(sc),
			)
			member.POST(
				"/score",
				middleware.TeacherValidationJSON(sc),
				classroomgin.UpdateStudentScore(sc),
			)
		}
	}

	lesson := authedRoutes.Group("/lesson")
	{
		lesson.POST(
			"/",
			lessongin.Create(sc),
		)
		lesson.PUT(
			"/",
			lessongin.Update(sc),
		)
		lesson.DELETE(
			"/",
			lessongin.Delete(sc),
		)
		lesson.GET("/", lessongin.List(sc))
		lesson.POST("/upload/:lesson_id", lessongin.UploadByFile(sc))
	}
}
