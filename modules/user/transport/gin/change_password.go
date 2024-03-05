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

func ChangePassword(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserChangePassword
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.DBMain).(*mongo.Client)

		store := userstorage.NewMgDBStorage(db)
		biz := userbiz.NewChangePasswordBiz(store)
		user := c.MustGet(appCommon.CurrentUser).(*usermodel.User)
		if err := biz.ChangePassword(c.Request.Context(), user, &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
