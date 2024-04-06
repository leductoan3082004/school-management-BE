package classroommodel

import (
	"SchoolManagement-BE/appCommon"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

const EntityName = "Classroom"

type MemberScore struct {
	Attendance uint `json:"attendance" bson:"attendance" `
	Lab        uint `json:"lab" bson:"lab" `
	Midterm    uint `json:"midterm" bson:"midterm"`
	Final      uint `json:"final" bson:"final"`
}
type Member struct {
	UserID      primitive.ObjectID `json:"user_id" bson:"user_id"`
	Role        int                `json:"role" bson:"role"`
	MemberScore `bson:",inline"`
}
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
	Members             []Member           `json:"members,omitempty" bson:"members"`
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

type ClassroomAddMember struct {
	ClassroomID string `json:"classroom_id" binding:"required"`
	UserID      string `json:"user_id" binding:"required"`
	Role        int    `json:"role" binding:"required"`
}

type ClassroomMemberList struct {
	ClassroomID string `form:"classroom_id" binding:"required"`
}

type ClassroomStudentScoreUpdate struct {
	ClassroomID string `json:"classroom_id" binding:"required"`
	UserID      string `json:"user_id" binding:"required"`
	Attendance  *uint  `json:"attendance"`
	Lab         *uint  `json:"lab"`
	Midterm     *uint  `json:"midterm"`
	Final       *uint  `json:"final"`
}

type ClassroomRemoveMember struct {
	ClassroomID string `json:"classroom_id" binding:"required"`
	UserID      string `json:"-"`
}

func (s *ClassroomStudentScoreUpdate) Validate() error {
	if s.Attendance != nil && *s.Attendance > 100 {
		return appCommon.ErrInvalidRequest(errors.New("attendance must be less than 100"))
	}
	if s.Lab != nil && *s.Lab > 100 {
		return appCommon.ErrInvalidRequest(errors.New("lab must be less than 100"))
	}
	if s.Midterm != nil && *s.Midterm > 100 {
		return appCommon.ErrInvalidRequest(errors.New("midterm must be less than 100"))
	}
	if s.Final != nil && *s.Final > 100 {
		return appCommon.ErrInvalidRequest(errors.New("final must be less than 100"))
	}
	return nil
}

var (
	ErrTeacherTimeTableOverlap = appCommon.NewCustomError(
		http.StatusBadRequest,
		errors.New("sum of ratio is not 100%"),
		"sum of ratio is not 100%",
		"ErrTeacherTimeTableOverlap",
	)
)
