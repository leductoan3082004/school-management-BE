package coursebiz

import (
	"SchoolManagement-BE/appCommon"
	coursemodel "SchoolManagement-BE/modules/course/model"
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
)

type courseListStore interface {
	ListDataWithCondition(ctx context.Context, paging *appCommon.Paging) ([]coursemodel.Course, error)
}

type courseListBiz struct {
	store  courseListStore
	logger logger.Logger
}

func NewCourseListBiz(store courseListStore) *courseListBiz {
	return &courseListBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("CourseListBiz"),
	}
}

func (biz *courseListBiz) ListDataWithCondition(ctx context.Context, paging *appCommon.Paging) ([]coursemodel.Course, error) {
	paging.Fulfill()
	res, err := biz.store.ListDataWithCondition(ctx, paging)
	if err != nil {
		biz.logger.WithSrc().Error(err)
		return []coursemodel.Course{}, appCommon.ErrCannotListEntity(coursemodel.EntityName, err)
	}
	return res, nil
}
