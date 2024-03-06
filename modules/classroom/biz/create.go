package classroombiz

import (
	"SchoolManagement-BE/appCommon"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	coursemodel "SchoolManagement-BE/modules/course/model"
	usermodel "SchoolManagement-BE/modules/user/model"
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type classroomCreateStore interface {
	Create(ctx context.Context, data *classroommodel.Classroom) error
}
type courseCheckingStore interface {
	FindById(ctx context.Context, id string) (*coursemodel.Course, error)
}

type userCheckingStore interface {
	FindById(ctx context.Context, id string) (*usermodel.User, error)
}

type createClassroomBiz struct {
	classStore  classroomCreateStore
	courseStore courseCheckingStore
	userStore   userCheckingStore
	logger      logger.Logger
}

func NewCreateClassroomBiz(
	classStore classroomCreateStore,
	courseStore courseCheckingStore,
	userStore userCheckingStore,
) *createClassroomBiz {
	return &createClassroomBiz{
		classStore:  classStore,
		courseStore: courseStore,
		userStore:   userStore,
		logger:      logger.GetCurrent().GetLogger("ClassroomCreateBiz"),
	}
}

func (biz *createClassroomBiz) CreateClassroom(
	ctx context.Context,
	data *classroommodel.ClassroomCreate,
) (*classroommodel.Classroom, error) {
	courseID, err := primitive.ObjectIDFromHex(data.CourseID)
	if err != nil {
		return nil, appCommon.ErrInvalidRequest(err)
	}
	teacherID, err := primitive.ObjectIDFromHex(data.TeacherID)
	if err != nil {
		return nil, appCommon.ErrInvalidRequest(err)
	}

	_, err = biz.courseStore.FindById(ctx, data.CourseID)
	if err != nil {
		if err == appCommon.ErrRecordNotFound {
			return nil, appCommon.ErrEntityNotFound(coursemodel.EntityName, err)
		}
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotGetEntity(coursemodel.EntityName, err)
	}

	_, err = biz.userStore.FindById(ctx, data.TeacherID)
	if err != nil {
		if err == appCommon.ErrRecordNotFound {
			return nil, appCommon.ErrEntityNotFound(usermodel.EntityName, err)
		}
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotGetEntity(usermodel.EntityName, err)
	}

	createData := &classroommodel.Classroom{
		CourseID:  courseID,
		TeacherID: teacherID,
		TimeTable: nil,
		Limit:     data.Limit,
	}

	timeTable := make([]classroommodel.TimeTable, 0)

	for i := range data.TimeTable {
		lessonStart := time.Unix(data.TimeTable[i].LessonStart, 0)
		lessonEnd := time.Unix(data.TimeTable[i].LessonEnd, 0)

		for j := 0; j < data.Weeks; j++ {
			if j > 0 {
				lessonStart = lessonStart.Add(7 * 24 * time.Hour)
				lessonEnd = lessonEnd.Add(7 * 24 * time.Hour)
			}
			timeTable = append(timeTable, classroommodel.TimeTable{
				ID:          primitive.NewObjectID(),
				LessonStart: lessonStart,
				LessonEnd:   lessonEnd,
			})
		}
	}
	createData.TimeTable = timeTable

	if err := biz.classStore.Create(ctx, createData); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotCreateEntity(classroommodel.EntityName, err)
	}

	return createData, nil
}
