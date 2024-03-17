package lessonbiz

import (
	"SchoolManagement-BE/appCommon"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	lessonmodel "SchoolManagement-BE/modules/lesson/model"
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
)

type listLessonStore interface {
	ListLesson(ctx context.Context, data *lessonmodel.LessonList, paging *appCommon.Paging) ([]lessonmodel.Lesson, error)
}
type classLessonListStore interface {
	FindById(ctx context.Context, id string) (*classroommodel.Classroom, error)
}
type listLessonBiz struct {
	store      listLessonStore
	classStore classLessonListStore
	logger     logger.Logger
}

func NewListLessonBiz(
	store listLessonStore,
	classStore classLessonListStore,
) *listLessonBiz {
	return &listLessonBiz{
		store:      store,
		classStore: classStore,
		logger:     logger.GetCurrent().GetLogger("ListLessonBiz"),
	}
}

func (biz *listLessonBiz) ListLesson(
	ctx context.Context,
	data *lessonmodel.LessonList,
	paging *appCommon.Paging,
) ([]lessonmodel.Lesson, error) {

	if paging == nil {
		paging = &appCommon.Paging{Page: 1, Limit: 10}
	}
	paging.Fulfill()

	result, err := biz.store.ListLesson(ctx, data, paging)
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotListEntity(lessonmodel.EntityName, err)
	}
	return result, nil
}
