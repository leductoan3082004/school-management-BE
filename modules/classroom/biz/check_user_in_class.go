package classroombiz

import (
	"SchoolManagement-BE/appCommon"
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
)

type checkUserInClassStore interface {
	CheckMemberInClass(
		ctx context.Context,
		classID string,
		userID string,
		role *int,
	) (bool, error)
}

type checkUserInClassBiz struct {
	store  checkUserInClassStore
	logger logger.Logger
}

func NewCheckUserInClassBiz(store checkUserInClassStore) *checkUserInClassBiz {
	return &checkUserInClassBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("CheckUserInClassBiz"),
	}
}

func (biz *checkUserInClassBiz) CheckUserInClass(
	ctx context.Context,
	classID string,
	userID string,
	role *int,
) (bool, error) {
	ok, err := biz.store.CheckMemberInClass(ctx, classID, userID, role)
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return false, appCommon.ErrInternal(err)
	}

	return ok, nil
}
