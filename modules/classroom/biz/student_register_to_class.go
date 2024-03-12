package classroombiz

import (
	"SchoolManagement-BE/appCommon"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	usermodel "SchoolManagement-BE/modules/user/model"
	"context"
	"errors"
	"github.com/lequocbinh04/go-sdk/logger"
)

type studentRegisterToClassStore interface {
	AddMemberToClass(ctx context.Context, data *classroommodel.Member, classID string) error
	CheckMemberInClass(
		ctx context.Context,
		classID string,
		userID string,
		role *int,
	) (bool, error)
}

type studentRegisterToClassBiz struct {
	store  studentRegisterToClassStore
	logger logger.Logger
}

func NewStudentRegisterToClassBiz(store studentRegisterToClassStore) *studentRegisterToClassBiz {
	return &studentRegisterToClassBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("StudentRegisterClassBiz"),
	}
}

func (biz *studentRegisterToClassBiz) AddMemberToClass(
	ctx context.Context,
	user *usermodel.User,
	classID string,
) error {
	exist, err := biz.store.CheckMemberInClass(ctx, classID, user.Id.Hex(), nil)
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrInternal(err)
	}
	if exist {
		return appCommon.ErrEntityExisted(usermodel.EntityName, nil)
	}

	if user.Role != usermodel.RoleUser {
		return appCommon.ErrInvalidRequest(errors.New("user must be student"))
	}

	member := classroommodel.Member{
		UserID: user.Id,
		Role:   user.Role,
	}

	if err := biz.store.AddMemberToClass(ctx, &member, classID); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotUpdateEntity(classroommodel.EntityName, err)
	}

	return nil
}
