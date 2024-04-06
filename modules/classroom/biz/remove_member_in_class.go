package classroombiz

import (
	"SchoolManagement-BE/appCommon"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
)

type removeMemberInClassStore interface {
	RemoveMemberInClass(ctx context.Context, classID string, userID string) error
}
type removeMemberInClassBiz struct {
	store  removeMemberInClassStore
	logger logger.Logger
}

func NewRemoveMemberInClassBiz(store removeMemberInClassStore) *removeMemberInClassBiz {
	return &removeMemberInClassBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("RemoveMemberInClassBiz"),
	}
}

func (biz *removeMemberInClassBiz) RemoveMemberInClass(ctx context.Context, data *classroommodel.ClassroomRemoveMember) error {
	err := biz.store.RemoveMemberInClass(ctx, data.ClassroomID, data.UserID)
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotUpdateEntity(classroommodel.EntityName, err)
	}
	return nil
}
