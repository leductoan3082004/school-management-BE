package middleware

import (
	"SchoolManagement-BE/appCommon"
	classroombiz "SchoolManagement-BE/modules/classroom/biz"
	classroomstorage "SchoolManagement-BE/modules/classroom/storage"
	lessonbiz "SchoolManagement-BE/modules/lesson/biz"
	lessonstorage "SchoolManagement-BE/modules/lesson/storage"
	usermodel "SchoolManagement-BE/modules/user/model"
	"errors"
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"go.mongodb.org/mongo-driver/mongo"
)

func ValidateGetMaterial(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(appCommon.CurrentUser).(*usermodel.User)

		if user.Role == usermodel.RoleAdmin {
			c.Next()
			return
		}

		key := c.Param("key")
		key = key[1:]

		db := sc.MustGet(appCommon.DBMain).(*mongo.Client)
		store := classroomstorage.NewMgDBStorage(db)
		lessonStore := lessonstorage.NewMgDBStorage(db)
		biz := classroombiz.NewCheckUserInClassBiz(store)
		lessonBiz := lessonbiz.NewFindLessonByKeyBiz(lessonStore)

		lesson, err := lessonBiz.FindLessonByKey(c.Request.Context(), key)
		if err != nil {
			panic(err)
		}

		ok, err := biz.CheckUserInClass(
			c.Request.Context(),
			lesson.ClassroomID.Hex(),
			user.Id.Hex(),
			nil,
		)

		if err != nil {
			panic(err)
		}

		if !ok {
			panic(appCommon.ErrNoPermission(errors.New("You are not a teacher in this classroom")))
		}

		c.Next()
	}
}
