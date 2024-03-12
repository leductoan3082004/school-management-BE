package classroommodel

import (
	"SchoolManagement-BE/appCommon"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

const EntityName = "Classroom"

type TimeTable struct {
	ID          primitive.ObjectID `json:"id" bson:"id"`
	LessonStart time.Time          `json:"lesson_start" bson:"lesson_start"`
	LessonEnd   time.Time          `json:"lesson_end" bson:"lesson_end"`
}

type TimeTables []TimeTable
type Classroom struct {
	appCommon.MgDBModel `json:",inline" bson:",inline"`
	CourseID            primitive.ObjectID `json:"course_id" bson:"course_id"`
	TimeTable           TimeTables         `json:"time_table" bson:"time_table"`
	Limit               int                `json:"limit" bson:"limit"`
}

func (s *TimeTables) CheckIntersect(other *TimeTables) bool {
	for _, t := range *s {
		for _, c := range *other {
			st := t.LessonStart
			if st.Before(c.LessonStart) {
				st = c.LessonStart
			}

			et := t.LessonEnd
			if et.After(c.LessonEnd) {
				et = c.LessonEnd
			}

			if st.Before(et) || st.Equal(et) {
				return true
			}
		}
	}
	return false

}
func (Classroom) TableName() string {
	return "classroom"
}

type ClassroomCreate struct {
	CourseID  string `json:"course_id" binding:"required"`
	Weeks     int    `json:"weeks" binding:"required"`
	Limit     int    `json:"limit" binding:"required"`
	TimeTable []struct {
		LessonStart int64 `json:"lesson_start" binding:"required"`
		LessonEnd   int64 `json:"lesson_end" binding:"required"`
	} `json:"time_table" binding:"required"`
}

func (data *ClassroomCreate) Validate() error {
	if len(data.TimeTable) == 0 {
		return appCommon.ErrInvalidRequest(errors.New("time_table is required"))
	}
	for i := range data.TimeTable {
		if data.TimeTable[i].LessonStart >= data.TimeTable[i].LessonEnd {
			return appCommon.ErrInvalidRequest(errors.New("lesson_start must be less than lesson_end"))
		}
	}
	return nil
}

type ClassroomUpdate struct {
	ClassroomId string    `json:"classroom_id" binding:"required"`
	Limit       *int      `json:"limit"`
	TimeIds     *[]string `json:"time_ids"`
}

type ClassroomDelete struct {
	ClassroomIds []string `json:"classroom_ids" binding:"required"`
}

type ClassroomList struct {
	CourseID string `form:"course_id" binding:"required"`
}

var (
	ErrTeacherTimeTableOverlap = appCommon.NewCustomError(
		http.StatusBadRequest,
		errors.New("sum of ratio is not 100%"),
		"sum of ratio is not 100%",
		"ErrTeacherTimeTableOverlap",
	)
)
