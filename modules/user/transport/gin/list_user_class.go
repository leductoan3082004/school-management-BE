package usergin

import (
	"SchoolManagement-BE/appCommon"
	classroomstorage "SchoolManagement-BE/modules/classroom/storage"
	coursestorage "SchoolManagement-BE/modules/course/storage"
	userbiz "SchoolManagement-BE/modules/user/biz"
	usermodel "SchoolManagement-BE/modules/user/model"
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func ListUserClass(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging appCommon.Paging
		var data usermodel.UserClassList
		if err := c.ShouldBind(&paging); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}

		db := sc.MustGet(appCommon.DBMain).(*mongo.Client)
		classStore := classroomstorage.NewMgDBStorage(db)
		courseStore := coursestorage.NewMgDBStorage(db)

		userID, err := primitive.ObjectIDFromHex(data.UserId)
		if err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}

		user := &usermodel.User{
			MgDBModel: appCommon.MgDBModel{
				Id: userID,
			},
		}
		biz := userbiz.NewListJoinedClassBiz(classStore, courseStore)

		res, err := biz.ListJoinedClass(c.Request.Context(), &paging, user)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, appCommon.NewSuccessResponse(res, paging, nil))
	}
}
