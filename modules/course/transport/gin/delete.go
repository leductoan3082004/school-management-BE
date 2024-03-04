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

func Delete(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data coursemodel.CourseDelete
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.DBMain).(*mongo.Client)

		store := coursestorage.NewMgDBStorage(db)
		biz := coursebiz.NewDeleteCourseBiz(store)
		if err := biz.DeleteData(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
