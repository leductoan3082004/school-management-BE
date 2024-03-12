package classroomgin

import (
	"SchoolManagement-BE/appCommon"
	classroombiz "SchoolManagement-BE/modules/classroom/biz"
	classroomstorage "SchoolManagement-BE/modules/classroom/storage"
	usermodel "SchoolManagement-BE/modules/user/model"
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func StudentRegisterClass(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		classID := c.Param("class_id")

		db := sc.MustGet(appCommon.DBMain).(*mongo.Client)
		user := c.MustGet(appCommon.CurrentUser).(*usermodel.User)
		classStore := classroomstorage.NewMgDBStorage(db)
		biz := classroombiz.NewStudentRegisterToClassBiz(classStore)
		if err := biz.AddMemberToClass(c.Request.Context(), user, classID); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
