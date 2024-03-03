package usergin

import (
	"SchoolManagement-BE/appCommon"
	usermodel "SchoolManagement-BE/modules/user/model"
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"net/http"
)

func GetProfile(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(appCommon.CurrentUser).(*usermodel.User)
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(u))
	}
}
