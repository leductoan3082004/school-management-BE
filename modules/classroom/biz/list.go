package classroombiz

import (
	"SchoolManagement-BE/appCommon"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
)

type classroomListStore interface {
	ListClassroom(ctx context.Context, paging *appCommon.Paging, filter *classroommodel.ClassroomList) ([]classroommodel.Classroom, error)
}

type classroomListBiz struct {
	store  classroomListStore
	logger logger.Logger
}

func NewClassroomListBiz(store classroomListStore) *classroomListBiz {
	return &classroomListBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("ClassroomListBiz"),
	}
}

func (biz *classroomListBiz) ListClassroom(
	ctx context.Context,
	paging *appCommon.Paging,
	filter *classroommodel.ClassroomList,
) ([]classroommodel.Classroom, error) {
	res, err := biz.store.ListClassroom(ctx, paging, filter)
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return []classroommodel.Classroom{}, appCommon.ErrCannotListEntity(classroommodel.EntityName, err)
	}
	return res, nil
}
