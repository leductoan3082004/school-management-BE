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

func Update(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data coursemodel.CourseUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.DBMain).(*mongo.Client)

		store := coursestorage.NewMgDBStorage(db)
		biz := coursebiz.NewCourseUpdateBiz(store)
		if err := biz.UpdateCourse(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
