package userbiz

import (
	"SchoolManagement-BE/appCommon"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	coursemodel "SchoolManagement-BE/modules/course/model"
	usermodel "SchoolManagement-BE/modules/user/model"
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type listJoinedClassStore interface {
	ListClassroomByUserID(
		ctx context.Context,
		paging *appCommon.Paging,
		userID string,
	) ([]classroommodel.Classroom, error)
}

type listCourseByClassIDsStore interface {
	ListCourseByIDs(ctx context.Context, CourseIds []string) ([]coursemodel.Course, error)
}

type listJoinedClassBiz struct {
	classStore  listJoinedClassStore
	courseStore listCourseByClassIDsStore
	logger      logger.Logger
}

func NewListJoinedClassBiz(
	classStore listJoinedClassStore,
	courseStore listCourseByClassIDsStore,
) *listJoinedClassBiz {
	return &listJoinedClassBiz{
		classStore:  classStore,
		courseStore: courseStore,
		logger:      logger.GetCurrent().GetLogger("ListJoinedClassBiz"),
	}
}

type response struct {
	Class  classroommodel.Classroom `json:"class"`
	Course coursemodel.Course       `json:"course"`
	Member classroommodel.Member    `json:"member"`
}

func (biz *listJoinedClassBiz) ListJoinedClass(
	ctx context.Context,
	paging *appCommon.Paging,
	user *usermodel.User,
) ([]response, error) {
	paging.Fulfill()
	classes, err := biz.classStore.ListClassroomByUserID(ctx, paging, user.Id.Hex())
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotListEntity(classroommodel.EntityName, err)
	}

	courseIds := make([]string, len(classes))
	for i, class := range classes {
		courseIds[i] = class.CourseID.Hex()
	}

	courses, err := biz.courseStore.ListCourseByIDs(ctx, courseIds)
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotListEntity(coursemodel.EntityName, err)
	}

	mp := make(map[primitive.ObjectID]coursemodel.Course)

	for _, course := range courses {
		mp[course.Id] = course
	}

	var res []response

	for _, class := range classes {
		temp := response{
			Class: class,
		}
		temp.Class.Members = []classroommodel.Member{}
		for _, member := range class.Members {
			if member.UserID == user.Id {
				temp.Member = member
				break
			}
		}
		temp.Course = mp[class.CourseID]
		res = append(res, temp)
	}

	if res == nil {
		return []response{}, nil
	}
	return res, nil
}
