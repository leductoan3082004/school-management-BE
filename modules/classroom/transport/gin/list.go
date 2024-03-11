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

func List(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data classroommodel.ClassroomList
		var paging appCommon.Paging
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		if err := c.ShouldBind(&paging); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}

		db := sc.MustGet(appCommon.DBMain).(*mongo.Client)

		classStore := classroomstorage.NewMgDBStorage(db)
		biz := classroombiz.NewClassroomListBiz(classStore)
		res, err := biz.ListClassroom(c.Request.Context(), &paging, &data)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.NewSuccessResponse(res, paging, nil))

	}
}
