package lessonbiz

import (
	"SchoolManagement-BE/appCommon"
	lessonmodel "SchoolManagement-BE/modules/lesson/model"
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
)

type findLessonByKeyStore interface {
	FindLessonByKey(ctx context.Context, key string) (*lessonmodel.Lesson, error)
}

type findLessonByKeyBiz struct {
	store  findLessonByKeyStore
	logger logger.Logger
}

func NewFindLessonByKeyBiz(store findLessonByKeyStore) *findLessonByKeyBiz {
	return &findLessonByKeyBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("FindLessonByKeyBiz"),
	}
}

func (biz *findLessonByKeyBiz) FindLessonByKey(
	ctx context.Context,
	key string,
) (*lessonmodel.Lesson, error) {
	lesson, err := biz.store.FindLessonByKey(ctx, key)
	if err != nil {
		if err == appCommon.ErrRecordNotFound {
			return nil, appCommon.ErrEntityNotFound(lessonmodel.EntityName, err)
		}
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotGetEntity(lessonmodel.EntityName, err)
	}

	return lesson, nil
}
