package middleware

import (
	"SchoolManagement-BE/appCommon"
	usermodel "SchoolManagement-BE/modules/user/model"
	"github.com/gin-gonic/gin"
)

func AdminAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(appCommon.CurrentUser).(*usermodel.User)
		if *user.Role != usermodel.RoleAdmin {
			panic(appCommon.ErrNoPermission(nil))
		}
		c.Next()
	}
}
