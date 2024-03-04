package coursebiz

import (
	"SchoolManagement-BE/appCommon"
	coursemodel "SchoolManagement-BE/modules/course/model"
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
	"time"
)

type courseCreateStore interface {
	Create(ctx context.Context, data *coursemodel.Course) error
}

type courseCreateBiz struct {
	store  courseCreateStore
	logger logger.Logger
}

func NewCourseCreateBiz(store courseCreateStore) *courseCreateBiz {
	return &courseCreateBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("CourseCreateBiz"),
	}
}

func (biz *courseCreateBiz) CreateCourse(ctx context.Context, data *coursemodel.CourseCreate) (*coursemodel.Course, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}
	createData := &coursemodel.Course{
		CourseSpec: coursemodel.CourseSpec{
			Limit:       int(data.Limit),
			CourseName:  data.CourseName,
			Credit:      data.Credit,
			Description: data.Description,
		},
		CourseRatio: coursemodel.CourseRatio{
			AttendanceRatio: int(data.AttendanceRatio),
			LabRatio:        int(data.LabRatio),
			MidtermRatio:    int(data.MidtermRatio),
			FinalRatio:      int(data.FinalRatio),
		},
		CourseRegisterTimeline: coursemodel.CourseRegisterTimeline{
			StartTime: time.Unix(data.StartTime, 0),
			EndTime:   time.Unix(data.EndTime, 0),
		},
	}

	if err := biz.store.Create(ctx, createData); err != nil {
		biz.logger.WithSrc().Error(err)
		return nil, appCommon.ErrCannotCreateEntity(coursemodel.EntityName, err)
	}
	return createData, nil
}
