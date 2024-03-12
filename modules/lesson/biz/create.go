package lessonbiz

import (
	"SchoolManagement-BE/appCommon"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	coursemodel "SchoolManagement-BE/modules/course/model"
	lessonmodel "SchoolManagement-BE/modules/lesson/model"
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
)

type lessonCreateStore interface {
	Create(ctx context.Context, data *lessonmodel.Lesson) error
}
type courseCreateStore interface {
	FindById(ctx context.Context, id string) (*coursemodel.Course, error)
}

type classCreateStore interface {
	FindById(ctx context.Context, id string) (*classroommodel.Classroom, error)
}

type createLessonBiz struct {
	store       lessonCreateStore
	courseStore courseCreateStore
	classStore  classCreateStore
	logger      logger.Logger
}

func NewCreateLessonBiz(
	store lessonCreateStore,
	courseStore courseCreateStore,
	classStore classCreateStore,
) *createLessonBiz {
	return &createLessonBiz{
		store:       store,
		courseStore: courseStore,
		classStore:  classStore,
		logger:      logger.GetCurrent().GetLogger("LessonCreateBiz"),
	}
}

func (biz *createLessonBiz) CreateLesson(ctx context.Context, data *lessonmodel.LessonCreate) error {
	lesson := lessonmodel.Lesson{
		Name:    data.Name,
		Content: data.Content,
	}

	// Check course
	if data.CourseID != nil {
		course, err := biz.courseStore.FindById(ctx, *data.CourseID)
		lesson.CourseID = course.Id
		if err != nil {
			if err == appCommon.ErrRecordNotFound {
				return appCommon.ErrEntityNotFound(coursemodel.EntityName, err)
			}
			biz.logger.WithSrc().Errorln(err)
			return appCommon.ErrCannotGetEntity(coursemodel.EntityName, err)
		}
	}

	// Check class
	if data.ClassID != nil {
		class, err := biz.classStore.FindById(ctx, *data.ClassID)
		lesson.ClassID = class.Id
		if err != nil {
			if err == appCommon.ErrRecordNotFound {
				return appCommon.ErrEntityNotFound(classroommodel.EntityName, err)
			}
			biz.logger.WithSrc().Errorln(err)
			return appCommon.ErrCannotGetEntity(classroommodel.EntityName, err)
		}
	}

	// Create lesson

	if err := biz.store.Create(ctx, &lesson); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotCreateEntity(lessonmodel.EntityName, err)
	}

	return nil
}