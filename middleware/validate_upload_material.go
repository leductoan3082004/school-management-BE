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

func TeacherUploadMaterialValidation(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(appCommon.CurrentUser).(*usermodel.User)

		if user.Role == usermodel.RoleAdmin {
			c.Next()
			return
		}

		lessonId := c.Param("lesson_id")

		db := sc.MustGet(appCommon.DBMain).(*mongo.Client)
		store := classroomstorage.NewMgDBStorage(db)
		lessonStore := lessonstorage.NewMgDBStorage(db)
		biz := classroombiz.NewCheckUserInClassBiz(store)
		lessonBiz := lessonbiz.NewLessonFindBiz(lessonStore)

		lesson, err := lessonBiz.FindLesson(c.Request.Context(), lessonId)
		if err != nil {
			panic(err)
		}

		teacher := usermodel.RoleTeacher
		ok, err := biz.CheckUserInClass(
			c.Request.Context(),
			lesson.ClassroomID.Hex(),
			user.Id.Hex(),
			&teacher,
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
