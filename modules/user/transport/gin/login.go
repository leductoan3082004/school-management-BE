package usergin

import (
	"SchoolManagement-BE/appCommon"
	userbiz "SchoolManagement-BE/modules/user/biz"
	usermodel "SchoolManagement-BE/modules/user/model"
	userstorage "SchoolManagement-BE/modules/user/storage"
	"SchoolManagement-BE/plugin/tokenprovider"
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func Login(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserLogin
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.DBMain).(*mongo.Client)
		tokenProvider := sc.MustGet(appCommon.PluginJWT).(tokenprovider.Provider)
		store := userstorage.NewMgDBStorage(db)
		biz := userbiz.NewUserLoginBiz(store, tokenProvider)
		token, err := biz.Login(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(token))
	}
}
