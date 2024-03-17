package classroombiz

import (
	"SchoolManagement-BE/appCommon"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	usermodel "SchoolManagement-BE/modules/user/model"
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
)

type addMemberToClassStore interface {
	AddMemberToClass(
		ctx context.Context,
		data *classroommodel.Member,
		classID string,
	) error
	CheckMemberInClass(
		ctx context.Context,
		classID string,
		userID string,
		role *int,
	) (bool, error)
	DecreaseLimit(ctx context.Context, classID string) error
}
type addMemberToClassUserStore interface {
	FindById(ctx context.Context, id string) (*usermodel.User, error)
}

type addMemberToClassBiz struct {
	store     addMemberToClassStore
	userStore addMemberToClassUserStore
	logger    logger.Logger
}

func NewAddMemberToClassBiz(
	store addMemberToClassStore,
	userStore addMemberToClassUserStore,
) *addMemberToClassBiz {
	return &addMemberToClassBiz{
		store:     store,
		userStore: userStore,
		logger:    logger.GetCurrent().GetLogger("AddMemberToClassBiz"),
	}
}

func (biz *addMemberToClassBiz) AddMemberToClass(
	ctx context.Context,
	data *classroommodel.ClassroomAddMember,
) error {

	user, err := biz.userStore.FindById(ctx, data.UserID)
	if err != nil {
		if err == appCommon.ErrRecordNotFound {
			return appCommon.ErrEntityNotFound(usermodel.EntityName, err)
		}
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotGetEntity(usermodel.EntityName, err)
	}

	isExist, err := biz.store.CheckMemberInClass(ctx, data.ClassroomID, data.UserID, nil)
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrInternal(err)
	}

	if isExist {
		return appCommon.ErrEntityExisted(usermodel.EntityName, nil)
	}

	member := classroommodel.Member{
		UserID: user.Id,
		Role:   data.Role,
	}

	if err := biz.store.AddMemberToClass(ctx, &member, data.ClassroomID); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotUpdateEntity(classroommodel.EntityName, err)
	}

	if err := biz.store.DecreaseLimit(ctx, data.ClassroomID); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotUpdateEntity(classroommodel.EntityName, err)
	}

	return nil
}
