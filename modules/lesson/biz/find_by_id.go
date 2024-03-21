package lessonbiz

import (
	"SchoolManagement-BE/appCommon"
	lessonmodel "SchoolManagement-BE/modules/lesson/model"
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
)

type lessonFindStore interface {
	Find(ctx context.Context, lessonID string) (*lessonmodel.Lesson, error)
}

type lessonFindBiz struct {
	store  lessonFindStore
	logger logger.Logger
}

func NewLessonFindBiz(store lessonFindStore) *lessonFindBiz {
	return &lessonFindBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("LessonFindBiz"),
	}
}

func (biz *lessonFindBiz) FindLesson(
	ctx context.Context,
	lessonID string,
) (*lessonmodel.Lesson, error) {
	lesson, err := biz.store.Find(ctx, lessonID)
	if err != nil {
		if err == appCommon.ErrRecordNotFound {
			return nil, appCommon.ErrEntityNotFound(lessonmodel.EntityName, err)
		}
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotGetEntity(lessonmodel.EntityName, err)
	}

	return lesson, nil
}
