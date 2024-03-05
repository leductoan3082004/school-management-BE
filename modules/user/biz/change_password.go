package userbiz

import (
	"SchoolManagement-BE/appCommon"
	usermodel "SchoolManagement-BE/modules/user/model"
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
)

type changePasswordStore interface {
	UpdatePassword(ctx context.Context, id, password string) error
}

type ChangePasswordBiz struct {
	store  changePasswordStore
	logger logger.Logger
}

func NewChangePasswordBiz(store changePasswordStore) *ChangePasswordBiz {
	return &ChangePasswordBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("ChangePasswordUserBiz"),
	}
}

func (biz *ChangePasswordBiz) ChangePassword(
	ctx context.Context,
	user *usermodel.User,
	data *usermodel.UserChangePassword,
) error {
	newPassword, err := appCommon.HMACEncode(data.NewPassword, user.Salt)
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrInternal(err)
	}
	oldPassword, err := appCommon.HMACEncode(data.OldPassword, user.Salt)
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrInternal(err)
	}

	if oldPassword != user.Password {
		return usermodel.ErrUsernameOrPasswordInvalid
	}

	if err := biz.store.UpdatePassword(ctx, user.Id.Hex(), newPassword); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}
	return nil
}
