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
	Count(ctx context.Context, courseID string) (int64, error)
	GetTeacherTimeTable(ctx context.Context, teacherID string) (classroommodel.TimeTables, error)
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

	if err := data.Validate(); err != nil {
		return nil, err
	}

	courseID, err := primitive.ObjectIDFromHex(data.CourseID)
	if err != nil {
		return nil, appCommon.ErrInvalidRequest(err)
	}
	teacherID, err := primitive.ObjectIDFromHex(data.TeacherID)
	if err != nil {
		return nil, appCommon.ErrInvalidRequest(err)
	}

	course, err := biz.courseStore.FindById(ctx, data.CourseID)
	if err != nil {
		if err == appCommon.ErrRecordNotFound {
			return nil, appCommon.ErrEntityNotFound(coursemodel.EntityName, err)
		}
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotGetEntity(coursemodel.EntityName, err)
	}

	count, err := biz.classStore.Count(ctx, data.CourseID)
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrInternal(err)
	}

	// check class exceed limit or not
	if count >= int64(course.Limit) {
		return nil, coursemodel.ErrClassroomLimitExceed
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

	// duplicate
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

	// check teacher timetable overlap
	teacherTimeTable, err := biz.classStore.GetTeacherTimeTable(ctx, data.TeacherID)
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotGetEntity(classroommodel.EntityName, err)
	}

	if createData.TimeTable.CheckIntersect(&teacherTimeTable) {
		return nil, classroommodel.ErrTeacherTimeTableOverlap
	}

	if err := biz.classStore.Create(ctx, createData); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotCreateEntity(classroommodel.EntityName, err)
	}

	return createData, nil
}
