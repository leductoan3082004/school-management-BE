package middleware

import (
	"SchoolManagement-BE/appCommon"
	classroombiz "SchoolManagement-BE/modules/classroom/biz"
	classroomstorage "SchoolManagement-BE/modules/classroom/storage"
	usermodel "SchoolManagement-BE/modules/user/model"
	"errors"
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"go.mongodb.org/mongo-driver/mongo"
)

func TeacherValidationForm(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data teacherValidation
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		user := c.MustGet(appCommon.CurrentUser).(*usermodel.User)

		if user.Role == usermodel.RoleAdmin {
			c.Next()
			return
		}

		db := sc.MustGet(appCommon.DBMain).(*mongo.Client)
		store := classroomstorage.NewMgDBStorage(db)
		biz := classroombiz.NewCheckUserInClassBiz(store)

		teacher := usermodel.RoleTeacher
		ok, err := biz.CheckUserInClass(
			c.Request.Context(),
			data.ClassroomID,
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
