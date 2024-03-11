package classroombiz

import (
	"SchoolManagement-BE/appCommon"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
)

type deleteClassroomStore interface {
	Delete(ctx context.Context, ids []string) error
}

type deleteClassroomBiz struct {
	store  deleteClassroomStore
	logger logger.Logger
}

func NewDeleteClassroomBiz(store deleteClassroomStore) *deleteClassroomBiz {
	return &deleteClassroomBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("ClassroomDeleteBiz"),
	}
}

func (biz *deleteClassroomBiz) DeleteClassroom(ctx context.Context, filter *classroommodel.ClassroomDelete) error {
	if err := biz.store.Delete(ctx, filter.ClassroomIds); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotDeleteEntity(classroommodel.EntityName, err)
	}
	return nil
}
