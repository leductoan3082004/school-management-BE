package middleware

import (
	"SchoolManagement-BE/appCommon"
	userbiz "SchoolManagement-BE/modules/user/biz"
	userstorage "SchoolManagement-BE/modules/user/storage"
	"SchoolManagement-BE/plugin/tokenprovider"
	"fmt"
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"github.com/lequocbinh04/go-sdk/sdkcm"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strings"
)

func ErrWrongAuthHeader(err error) *sdkcm.AppError {
	return appCommon.NewCustomError(
		http.StatusUnauthorized,
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	//"Authorization" : "Bearer {token}"

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

func RequiredAuth(sc goservice.ServiceContext) func(c *gin.Context) {
	tokenProvider := sc.MustGet(appCommon.PluginJWT).(tokenprovider.Provider)
	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}

		db := sc.MustGet(appCommon.DBMain).(*mongo.Client)
		store := userstorage.NewMgDBStorage(db)
		userBiz := userbiz.NewFindUserByIdBiz(store)
		user, err := userBiz.FindById(c.Request.Context(), payload.UserId)

		if err != nil {
			panic(err)
		}

		c.Set(appCommon.CurrentUser, user)
		c.Next()
	}
}
