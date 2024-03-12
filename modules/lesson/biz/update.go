package lessonbiz

import (
	"SchoolManagement-BE/appCommon"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	coursemodel "SchoolManagement-BE/modules/course/model"
	lessonmodel "SchoolManagement-BE/modules/lesson/model"
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
)

type lessonUpdateStore interface {
	Update(ctx context.Context, data *lessonmodel.LessonUpdate) error
}
type courseUpdateStore interface {
	FindById(ctx context.Context, id string) (*coursemodel.Course, error)
}

type classUpdateStore interface {
	FindById(ctx context.Context, id string) (*classroommodel.Classroom, error)
}

type updateLessonBiz struct {
	store       lessonUpdateStore
	courseStore courseUpdateStore
	classStore  classUpdateStore
	logger      logger.Logger
}

func NewUpdateLessonBiz(
	store lessonUpdateStore,
	courseStore courseUpdateStore,
	classStore classUpdateStore,
) *updateLessonBiz {
	return &updateLessonBiz{
		store:       store,
		courseStore: courseStore,
		classStore:  classStore,
		logger:      logger.GetCurrent().GetLogger("LessonUpdateBiz"),
	}
}

func (biz *updateLessonBiz) UpdateLesson(ctx context.Context, data *lessonmodel.LessonUpdate) error {
	// Check lesson
	if data.CourseID != nil {
		_, err := biz.courseStore.FindById(ctx, *data.CourseID)
		if err != nil {
			if err == appCommon.ErrRecordNotFound {
				return appCommon.ErrEntityNotFound(coursemodel.EntityName, err)
			}
			biz.logger.WithSrc().Errorln(err)
			return appCommon.ErrCannotGetEntity(coursemodel.EntityName, err)
		}
	}
	if data.ClassID != nil {
		_, err := biz.classStore.FindById(ctx, *data.ClassID)
		if err != nil {
			if err == appCommon.ErrRecordNotFound {
				return appCommon.ErrEntityNotFound(classroommodel.EntityName, err)
			}
			biz.logger.WithSrc().Errorln(err)
			return appCommon.ErrCannotGetEntity(classroommodel.EntityName, err)
		}
	}

	if err := biz.store.Update(ctx, data); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotUpdateEntity(lessonmodel.EntityName, err)
	}
	return nil
}
