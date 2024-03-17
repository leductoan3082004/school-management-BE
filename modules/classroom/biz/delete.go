package classroombiz

import (
	"SchoolManagement-BE/appCommon"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	lessonmodel "SchoolManagement-BE/modules/lesson/model"
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
)

type deleteClassroomStore interface {
	Delete(ctx context.Context, ids []string) error
}
type deleteLessonStore interface {
	DeleteLessonByClassID(ctx context.Context, classID []string) error
}
type deleteClassroomBiz struct {
	store       deleteClassroomStore
	lessonStore deleteLessonStore
	logger      logger.Logger
}

func NewDeleteClassroomBiz(
	store deleteClassroomStore,
	lessonStore deleteLessonStore,
) *deleteClassroomBiz {
	return &deleteClassroomBiz{
		store:       store,
		lessonStore: lessonStore,
		logger:      logger.GetCurrent().GetLogger("ClassroomDeleteBiz"),
	}
}

func (biz *deleteClassroomBiz) DeleteClassroom(ctx context.Context, filter *classroommodel.ClassroomDelete) error {
	if err := biz.store.Delete(ctx, filter.ClassroomIds); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotDeleteEntity(classroommodel.EntityName, err)
	}
	if err := biz.lessonStore.DeleteLessonByClassID(ctx, filter.ClassroomIds); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotDeleteEntity(lessonmodel.EntityName, err)
	}
	return nil
}
