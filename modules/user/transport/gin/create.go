package usergin

import (
	"SchoolManagement-BE/appCommon"
	userbiz "SchoolManagement-BE/modules/user/biz"
	usermodel "SchoolManagement-BE/modules/user/model"
	userstorage "SchoolManagement-BE/modules/user/storage"
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func Create(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.DBMain).(*mongo.Client)

		store := userstorage.NewMgDBStorage(db)
		biz := userbiz.NewUserCreateBiz(store)
		if err := biz.Create(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
