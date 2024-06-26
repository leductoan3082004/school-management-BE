package userbiz

import (
	"SchoolManagement-BE/appCommon"
	usermodel "SchoolManagement-BE/modules/user/model"
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
)

type findUserByIdStore interface {
	FindById(ctx context.Context, id string) (*usermodel.User, error)
}

type FindUserByIdBiz struct {
	store  findUserByIdStore
	logger logger.Logger
}

func NewFindUserByIdBiz(store findUserByIdStore) *FindUserByIdBiz {
	return &FindUserByIdBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("FindUserByIdBiz"),
	}
}

func (biz *FindUserByIdBiz) FindById(ctx context.Context, id string) (*usermodel.User, error) {
	user, err := biz.store.FindById(ctx, id)
	if err != nil {
		if err == appCommon.ErrRecordNotFound {
			return nil, appCommon.ErrEntityNotFound(usermodel.EntityName, err)
		}
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotGetEntity(usermodel.EntityName, err)
	}
	return user, nil
}
