package coursegin

import (
	"SchoolManagement-BE/appCommon"
	coursebiz "SchoolManagement-BE/modules/course/biz"
	coursemodel "SchoolManagement-BE/modules/course/model"
	coursestorage "SchoolManagement-BE/modules/course/storage"
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func List(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging appCommon.Paging
		var filter coursemodel.CourseList
		if err := c.ShouldBind(&paging); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		if err := c.ShouldBind(&filter); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.DBMain).(*mongo.Client)

		store := coursestorage.NewMgDBStorage(db)
		biz := coursebiz.NewCourseListBiz(store)
		res, err := biz.ListDataWithCondition(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.NewSuccessResponse(res, paging, nil))
	}
}
