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

type studentValidation struct {
	ClassroomID string `json:"classroom_id" form:"classroom_id" binding:"required"`
}

func StudentValidationForm(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data studentValidation
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

		ok, err := biz.CheckUserInClass(
			c.Request.Context(),
			data.ClassroomID,
			user.Id.Hex(),
			nil,
		)

		if err != nil {
			panic(err)
		}

		if !ok {
			panic(appCommon.ErrNoPermission(errors.New("You are not in this classroom")))
		}

		c.Next()
	}
}
