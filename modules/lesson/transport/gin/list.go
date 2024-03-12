package lessongin

import (
	"SchoolManagement-BE/appCommon"
	classroomstorage "SchoolManagement-BE/modules/classroom/storage"
	lessonbiz "SchoolManagement-BE/modules/lesson/biz"
	lessonmodel "SchoolManagement-BE/modules/lesson/model"
	lessonstorage "SchoolManagement-BE/modules/lesson/storage"
	usermodel "SchoolManagement-BE/modules/user/model"
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func List(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data lessonmodel.LessonList
		var paging appCommon.Paging

		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		if err := c.ShouldBind(&paging); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.DBMain).(*mongo.Client)

		store := lessonstorage.NewMgDBStorage(db)
		classStore := classroomstorage.NewMgDBStorage(db)
		user := c.MustGet(appCommon.CurrentUser).(*usermodel.User)
		biz := lessonbiz.NewListLessonBiz(store, classStore, nil)

		res, err := biz.ListLesson(c.Request.Context(), user, &data, &paging)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.NewSuccessResponse(res, paging, nil))
	}
}
