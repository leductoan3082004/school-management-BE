package lessonbiz

import (
	"SchoolManagement-BE/appCommon"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	lessonmodel "SchoolManagement-BE/modules/lesson/model"
	usermodel "SchoolManagement-BE/modules/user/model"
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
)

type listLessonStore interface {
	ListLesson(ctx context.Context, data *lessonmodel.LessonList, paging *appCommon.Paging) ([]lessonmodel.Lesson, error)
}
type classLessonListStore interface {
	FindById(ctx context.Context, id string) (*classroommodel.Classroom, error)
}
type classMemberListStore interface {
	FindByUserId(ctx context.Context, classID string, userID string) error
}
type listLessonBiz struct {
	store            listLessonStore
	classStore       classLessonListStore
	classMemberStore classMemberListStore
	logger           logger.Logger
}

func NewListLessonBiz(
	store listLessonStore,
	classStore classLessonListStore,
	classMemberStore classMemberListStore,
) *listLessonBiz {
	return &listLessonBiz{
		store:            store,
		classStore:       classStore,
		classMemberStore: classMemberStore,
		logger:           logger.GetCurrent().GetLogger("ListLessonBiz"),
	}
}

func (biz *listLessonBiz) ListLesson(
	ctx context.Context,
	user *usermodel.User,
	data *lessonmodel.LessonList,
	paging *appCommon.Paging,
) ([]lessonmodel.Lesson, error) {
	if user.Role != usermodel.RoleAdmin {
		err := biz.classMemberStore.FindByUserId(ctx, data.ClassID, user.Id.Hex())
		if err != nil {
			if err == appCommon.ErrRecordNotFound {
				return nil, appCommon.ErrNoPermission(nil)
			}
			biz.logger.WithSrc().Errorln(err)
			return nil, appCommon.ErrCannotGetEntity(lessonmodel.EntityName, err)
		}
	}

	if paging == nil {
		paging = &appCommon.Paging{Page: 1, Limit: 10}
	}
	paging.Fulfill()

	// check if student or teacher has permission to access the class or course

	result, err := biz.store.ListLesson(ctx, data, paging)
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotListEntity(lessonmodel.EntityName, err)
	}
	return result, nil
}
