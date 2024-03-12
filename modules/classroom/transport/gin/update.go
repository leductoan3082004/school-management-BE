package classroomgin

import (
	"SchoolManagement-BE/appCommon"
	classroombiz "SchoolManagement-BE/modules/classroom/biz"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	classroomstorage "SchoolManagement-BE/modules/classroom/storage"
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func Update(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data classroommodel.ClassroomUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.DBMain).(*mongo.Client)

		classStore := classroomstorage.NewMgDBStorage(db)
		biz := classroombiz.NewClassUpdateBiz(classStore)
		if err := biz.UpdateClass(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))

	}
}
