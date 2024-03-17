package classroombiz

import (
	"SchoolManagement-BE/appCommon"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
)

type classUpdateStore interface {
	UpdateClass(ctx context.Context, data *classroommodel.ClassroomUpdate) error
}

type classUpdateBiz struct {
	store  classUpdateStore
	logger logger.Logger
}

func NewClassUpdateBiz(store classUpdateStore) *classUpdateBiz {
	return &classUpdateBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("ClassUpdateBiz"),
	}
}

func (biz *classUpdateBiz) UpdateClass(ctx context.Context, data *classroommodel.ClassroomUpdate) error {
	if err := biz.store.UpdateClass(ctx, data); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotUpdateEntity(classroommodel.EntityName, err)
	}
	return nil
}
