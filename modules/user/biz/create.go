package userbiz

import (
	"SchoolManagement-BE/appCommon"
	usermodel "SchoolManagement-BE/modules/user/model"
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
)

type userCreateStore interface {
	Create(ctx context.Context, data *usermodel.User) error
	FindByUsername(ctx context.Context, username string) (*usermodel.User, error)
}

type userCreateBiz struct {
	store  userCreateStore
	logger logger.Logger
}

func NewUserCreateBiz(store userCreateStore) *userCreateBiz {
	return &userCreateBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("userBizRegister"),
	}
}

func (biz *userCreateBiz) Create(ctx context.Context, data *usermodel.UserCreate) error {
	_, err := biz.store.FindByUsername(ctx, data.Username)
	if err != nil {
		if err != appCommon.ErrRecordNotFound {
			biz.logger.WithSrc().Errorln(err)
			return appCommon.ErrInternal(err)
		}
	} else {
		return appCommon.ErrEntityExisted(usermodel.EntityName, nil)
	}

	salt := appCommon.GenSalt(30)
	password, err := appCommon.HMACEncode(data.Password, salt)
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrInternal(err)
	}

	user := &usermodel.User{
		AuthenticationUser: usermodel.AuthenticationUser{
			Username: data.Username,
			Password: password,
			Salt:     salt,
		},
		SpecUser: usermodel.SpecUser{
			Role:    data.Role,
			Name:    data.Name,
			Phone:   data.Phone,
			Address: data.Address,
		},
	}

	if err := biz.store.Create(ctx, user); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}
