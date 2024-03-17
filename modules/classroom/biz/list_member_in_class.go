package classroombiz

import (
	"SchoolManagement-BE/appCommon"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	usermodel "SchoolManagement-BE/modules/user/model"
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type listMemberInClassStore interface {
	FindById(ctx context.Context, id string) (*classroommodel.Classroom, error)
}

type listMemberInClassUserStore interface {
	FindMemberByIds(ctx context.Context, id []string) ([]usermodel.User, error)
}

type ListMemberInClassBiz struct {
	store     listMemberInClassStore
	userStore listMemberInClassUserStore
	logger    logger.Logger
}

func NewListMemberInClassBiz(
	store listMemberInClassStore,
	userStore listMemberInClassUserStore,
) *ListMemberInClassBiz {
	return &ListMemberInClassBiz{
		store:     store,
		userStore: userStore,
		logger:    logger.GetCurrent().GetLogger("ListMemberInClassroomBiz"),
	}
}

func (biz *ListMemberInClassBiz) ListMemberInClass(
	ctx context.Context,
	filter *classroommodel.ClassroomMemberList,
) ([]usermodel.User, error) {
	data, err := biz.store.FindById(ctx, filter.ClassroomID)

	if err != nil {
		if err == appCommon.ErrRecordNotFound {
			return []usermodel.User{}, appCommon.ErrEntityNotFound(classroommodel.EntityName, err)
		}
		biz.logger.WithSrc().Errorln(err)
		return []usermodel.User{}, appCommon.ErrCannotGetEntity(classroommodel.EntityName, err)
	}

	memberIds := make([]string, len(data.Members))

	mp := make(map[primitive.ObjectID]int)
	for i, v := range data.Members {
		memberIds[i] = v.UserID.Hex()
		mp[v.UserID] = v.Role
	}

	users, err := biz.userStore.FindMemberByIds(ctx, memberIds)
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return []usermodel.User{}, appCommon.ErrCannotListEntity(usermodel.EntityName, err)
	}

	for i, v := range users {
		users[i].Role = mp[v.Id]
	}

	return users, nil
}
