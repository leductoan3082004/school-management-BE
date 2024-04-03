package coursebiz

import (
	"SchoolManagement-BE/appCommon"
	coursemodel "SchoolManagement-BE/modules/course/model"
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
)

type findCourseStore interface {
	FindById(ctx context.Context, id string) (*coursemodel.Course, error)
}
type findCourseBiz struct {
	store  findCourseStore
	logger logger.Logger
}

func NewFindCourseBiz(store findCourseStore) *findCourseBiz {
	return &findCourseBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("FindCourseBiz"),
	}
}

func (biz *findCourseBiz) FindById(ctx context.Context, id string) (*coursemodel.Course, error) {
	data, err := biz.store.FindById(ctx, id)
	if err != nil {
		if err == appCommon.ErrRecordNotFound {
			return nil, appCommon.ErrEntityNotFound(coursemodel.EntityName, err)
		}
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotGetEntity(coursemodel.EntityName, err)
	}
	return data, nil
}
