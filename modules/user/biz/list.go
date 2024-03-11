package userbiz

import (
	"SchoolManagement-BE/appCommon"
	usermodel "SchoolManagement-BE/modules/user/model"
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
)

type listUserStore interface {
	ListDataWithCondition(ctx context.Context, filter *usermodel.UserList, paging *appCommon.Paging, moreInfo ...string) ([]usermodel.User, error)
}

type listUserBiz struct {
	store  listUserStore
	logger logger.Logger
}

func NewListUserBiz(store listUserStore) *listUserBiz {
	return &listUserBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("ListUserBiz"),
	}
}

func (biz *listUserBiz) ListUser(ctx context.Context, filter *usermodel.UserList, paging *appCommon.Paging) ([]usermodel.User, error) {
	res, err := biz.store.ListDataWithCondition(ctx, filter, paging)
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return []usermodel.User{}, appCommon.ErrCannotListEntity(usermodel.EntityName, err)
	}
	return res, nil
}
