package lessongin

import (
	"SchoolManagement-BE/appCommon"
	classroomstorage "SchoolManagement-BE/modules/classroom/storage"
	lessonbiz "SchoolManagement-BE/modules/lesson/biz"
	lessonmodel "SchoolManagement-BE/modules/lesson/model"
	lessonstorage "SchoolManagement-BE/modules/lesson/storage"
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func Create(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data lessonmodel.LessonCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.DBMain).(*mongo.Client)

		store := lessonstorage.NewMgDBStorage(db)
		classStore := classroomstorage.NewMgDBStorage(db)

		biz := lessonbiz.NewCreateLessonBiz(store, classStore)
		if err := biz.CreateLesson(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
