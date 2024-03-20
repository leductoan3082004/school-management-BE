package classroombiz

import (
	"SchoolManagement-BE/appCommon"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
)

type updateStudentScoreStore interface {
	UpdateStudentScore(
		ctx context.Context,
		data *classroommodel.ClassroomStudentScoreUpdate,
	) error
}

type updateStudentScoreBiz struct {
	store  updateStudentScoreStore
	logger logger.Logger
}

func NewUpdateStudentScoreBiz(store updateStudentScoreStore) *updateStudentScoreBiz {
	return &updateStudentScoreBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("ClassroomUpdateStudentScoreBiz"),
	}
}

func (biz *updateStudentScoreBiz) UpdateStudentScore(
	ctx context.Context,
	data *classroommodel.ClassroomStudentScoreUpdate,
) error {
	if err := data.Validate(); err != nil {
		return err
	}
	if err := biz.store.UpdateStudentScore(ctx, data); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotUpdateEntity(classroommodel.EntityName, err)
	}
	return nil
}
