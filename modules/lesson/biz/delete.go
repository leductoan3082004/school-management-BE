package lessonbiz

import (
	"SchoolManagement-BE/appCommon"
	lessonmodel "SchoolManagement-BE/modules/lesson/model"
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
)

type deleteLessonStore interface {
	DeleteLesson(ctx context.Context, data *lessonmodel.LessonDelete) error
}

type deleteLessonBiz struct {
	store  deleteLessonStore
	logger logger.Logger
}

func NewDeleteLessonBiz(store deleteLessonStore) *deleteLessonBiz {
	return &deleteLessonBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("DeleteLessonBiz"),
	}
}

func (biz *deleteLessonBiz) DeleteLesson(ctx context.Context, data *lessonmodel.LessonDelete) error {
	if err := biz.store.DeleteLesson(ctx, data); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotDeleteEntity(lessonmodel.EntityName, err)
	}
	return nil
}
