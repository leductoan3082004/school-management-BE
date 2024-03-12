package lessonbiz

import (
	"SchoolManagement-BE/appCommon"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	lessonmodel "SchoolManagement-BE/modules/lesson/model"
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
)

type lessonUpdateStore interface {
	Update(ctx context.Context, data *lessonmodel.LessonUpdate) error
}

type classUpdateStore interface {
	FindById(ctx context.Context, id string) (*classroommodel.Classroom, error)
}

type updateLessonBiz struct {
	store      lessonUpdateStore
	classStore classUpdateStore
	logger     logger.Logger
}

func NewUpdateLessonBiz(
	store lessonUpdateStore,
	classStore classUpdateStore,
) *updateLessonBiz {
	return &updateLessonBiz{
		store:      store,
		classStore: classStore,
		logger:     logger.GetCurrent().GetLogger("LessonUpdateBiz"),
	}
}

func (biz *updateLessonBiz) UpdateLesson(ctx context.Context, data *lessonmodel.LessonUpdate) error {
	_, err := biz.classStore.FindById(ctx, data.ClassID)
	if err != nil {
		if err == appCommon.ErrRecordNotFound {
			return appCommon.ErrEntityNotFound(classroommodel.EntityName, err)
		}
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotGetEntity(classroommodel.EntityName, err)
	}

	if err := biz.store.Update(ctx, data); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotUpdateEntity(lessonmodel.EntityName, err)
	}
	return nil
}
