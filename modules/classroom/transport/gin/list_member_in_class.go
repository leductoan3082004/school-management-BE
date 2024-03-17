package classroomgin

import (
	"SchoolManagement-BE/appCommon"
	classroombiz "SchoolManagement-BE/modules/classroom/biz"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	classroomstorage "SchoolManagement-BE/modules/classroom/storage"
	userstorage "SchoolManagement-BE/modules/user/storage"
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func ListMemberInClass(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data classroommodel.ClassroomMemberList
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}

		db := sc.MustGet(appCommon.DBMain).(*mongo.Client)

		classStore := classroomstorage.NewMgDBStorage(db)
		userStore := userstorage.NewMgDBStorage(db)
		biz := classroombiz.NewListMemberInClassBiz(classStore, userStore)
		res, err := biz.ListMemberInClass(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(res))

	}
}
