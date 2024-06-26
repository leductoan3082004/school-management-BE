package userbiz

import (
	"SchoolManagement-BE/appCommon"
	usermodel "SchoolManagement-BE/modules/user/model"
	"SchoolManagement-BE/plugin/tokenprovider"
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
)

type userLoginStore interface {
	FindByUsername(ctx context.Context, username string) (*usermodel.User, error)
}
type userLoginBiz struct {
	store         userLoginStore
	logger        logger.Logger
	tokenProvider tokenprovider.Provider
}

func NewUserLoginBiz(store userLoginStore, tokenProvider tokenprovider.Provider) *userLoginBiz {
	return &userLoginBiz{
		store:         store,
		tokenProvider: tokenProvider,
		logger:        logger.GetCurrent().GetLogger("UserLoginBiz"),
	}
}

func (biz *userLoginBiz) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	user, err := biz.store.FindByUsername(ctx, data.Username)
	if err != nil {
		if err == appCommon.ErrRecordNotFound {
			return nil, appCommon.ErrEntityNotFound(usermodel.EntityName, err)
		}
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotGetEntity(usermodel.EntityName, err)
	}

	password, err := appCommon.HMACEncode(data.Password, user.Salt)
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrInternal(err)
	}
	if password != user.Password {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	sessionId := appCommon.GenSalt(20)
	payload := tokenprovider.TokenPayload{
		UserId:    user.Id.Hex(),
		Role:      "user",
		SessionID: sessionId,
	}
	token, err := biz.tokenProvider.Generate(payload, appCommon.ExpiryAccessToken)
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrInternal(err)
	}
	return token, nil
}
