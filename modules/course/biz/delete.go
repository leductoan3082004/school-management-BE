package coursebiz

import (
	"SchoolManagement-BE/appCommon"
	coursemodel "SchoolManagement-BE/modules/course/model"
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
)

type deleteCourseStore interface {
	Delete(ctx context.Context, id string) error
}
type deleteCourseBiz struct {
	store  deleteCourseStore
	logger logger.Logger
}

func NewDeleteCourseBiz(store deleteCourseStore) *deleteCourseBiz {
	return &deleteCourseBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("DeleteCourseBiz"),
	}
}

func (biz *deleteCourseBiz) DeleteData(ctx context.Context, filter *coursemodel.CourseDelete) error {
	if err := biz.store.Delete(ctx, filter.CourseId); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotDeleteEntity(coursemodel.EntityName, err)
	}
	return nil
}
