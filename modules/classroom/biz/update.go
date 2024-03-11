package classroombiz

import (
	"SchoolManagement-BE/appCommon"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	usermodel "SchoolManagement-BE/modules/user/model"
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
)

type classUpdateStore interface {
	UpdateClass(ctx context.Context, data *classroommodel.ClassroomUpdate) error
}

type classUpdateBiz struct {
	store     classUpdateStore
	userStore userCheckingStore
	logger    logger.Logger
}

func NewClassUpdateBiz(store classUpdateStore, userStore userCheckingStore) *classUpdateBiz {
	return &classUpdateBiz{
		store:     store,
		userStore: userStore,
		logger:    logger.GetCurrent().GetLogger("ClassUpdateBiz"),
	}
}

func (biz *classUpdateBiz) UpdateClass(ctx context.Context, data *classroommodel.ClassroomUpdate) error {
	if data.TeacherID != nil {
		_, err := biz.userStore.FindById(ctx, *data.TeacherID)
		if err != nil {
			if err == appCommon.ErrRecordNotFound {
				return appCommon.ErrEntityNotFound(usermodel.EntityName, err)
			}
			biz.logger.WithSrc().Errorln(err)
			return appCommon.ErrCannotGetEntity(usermodel.EntityName, err)
		}
	}
	if err := biz.store.UpdateClass(ctx, data); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotUpdateEntity(classroommodel.EntityName, err)
	}
	return nil
}
