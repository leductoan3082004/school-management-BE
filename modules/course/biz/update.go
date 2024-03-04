package coursebiz

import (
	"SchoolManagement-BE/appCommon"
	coursemodel "SchoolManagement-BE/modules/course/model"
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
)

type courseUpdateStore interface {
	FindById(ctx context.Context, id string) (*coursemodel.Course, error)
	Update(ctx context.Context, updateData *coursemodel.CourseUpdate) error
}

type courseUpdateBiz struct {
	store  courseUpdateStore
	logger logger.Logger
}

func NewCourseUpdateBiz(store courseUpdateStore) *courseUpdateBiz {
	return &courseUpdateBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("CourseUpdateBiz"),
	}
}

func (biz *courseUpdateBiz) UpdateCourse(ctx context.Context, data *coursemodel.CourseUpdate) error {
	course, err := biz.store.FindById(ctx, data.CourseId)

	if err != nil {
		if err == appCommon.ErrRecordNotFound {
			return appCommon.ErrEntityNotFound(coursemodel.EntityName, err)
		}
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotGetEntity(coursemodel.EntityName, err)
	}

	if data.MidtermRatio != nil {
		course.MidtermRatio = int(*data.MidtermRatio)
	}
	if data.FinalRatio != nil {
		course.FinalRatio = int(*data.FinalRatio)
	}
	if data.LabRatio != nil {
		course.LabRatio = int(*data.LabRatio)
	}
	if data.AttendanceRatio != nil {
		course.AttendanceRatio = int(*data.AttendanceRatio)
	}

	if err := course.Validate(); err != nil {
		return appCommon.ErrInvalidRequest(err)
	}

	if err := biz.store.Update(ctx, data); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotUpdateEntity(coursemodel.EntityName, err)
	}

	return nil
}
