package classroomgin

import (
	"SchoolManagement-BE/appCommon"
	classroombiz "SchoolManagement-BE/modules/classroom/biz"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	classroomstorage "SchoolManagement-BE/modules/classroom/storage"
	usermodel "SchoolManagement-BE/modules/user/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	goservice "github.com/lequocbinh04/go-sdk"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func RemoveMemberInClass(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data classroommodel.ClassroomRemoveMember
		if err := c.ShouldBindBodyWith(&data, binding.JSON); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}

		db := sc.MustGet(appCommon.DBMain).(*mongo.Client)

		user := c.MustGet(appCommon.CurrentUser).(*usermodel.User)
		classStore := classroomstorage.NewMgDBStorage(db)
		biz := classroombiz.NewRemoveMemberInClassBiz(classStore)
		data.UserID = user.Id.Hex()
		if err := biz.RemoveMemberInClass(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))

	}
}
