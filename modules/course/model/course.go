package coursemodel

import (
	"SchoolManagement-BE/appCommon"
	"errors"
	"net/http"
)

const EntityName = "Course"

type CourseRatio struct {
	AttendanceRatio int `json:"attendance_ratio" bson:"attendance_ratio"`
	LabRatio        int `json:"lab_ratio" bson:"lab_ratio"`
	MidtermRatio    int `json:"midterm_ratio" bson:"midterm_ratio"`
	FinalRatio      int `json:"final_ratio" bson:"final_ratio"`
}
type CourseSpec struct {
	Limit      int    `json:"limit" bson:"limit"`
	CourseName string `json:"course_name" bson:"course_name"`
	Credit     int    `json:"credit" bson:"credit"`
}
type Course struct {
	appCommon.MgDBModel `json:",inline" bson:",inline"`
	CourseSpec          `json:",inline" bson:",inline"`
	CourseRatio         `json:",inline" bson:",inline"`
}

func (Course) TableName() string {
	return "course"
}

type CourseUpdate struct {
	AttendanceRatio *uint   `json:"attendance_ratio"`
	LabRatio        *uint   `json:"lab_ratio"`
	MidtermRatio    *uint   `json:"midterm_ratio"`
	FinalRatio      *uint   `json:"final_ratio"`
	Limit           *uint   `json:"limit"`
	CourseName      *string `json:"course_name"`
	Credit          *int    `json:"credit"`
}
type CourseCreate struct {
	AttendanceRatio uint   `json:"attendance_ratio"`
	LabRatio        uint   `json:"lab_ratio"`
	MidtermRatio    uint   `json:"midterm_ratio"`
	FinalRatio      uint   `json:"final_ratio"`
	Limit           uint   `json:"limit"`
	CourseName      string `json:"course_name" binding:"required"`
	Credit          int    `json:"credit"`
}

type CourseDelete struct {
	CourseId string `json:"course_id" binding:"required"`
}

func (s *CourseCreate) Validate() error {
	if s.FinalRatio+s.AttendanceRatio+s.MidtermRatio+s.LabRatio != 100 {
		return ErrInvalidRatio
	}
	return nil
}

var (
	ErrInvalidRatio = appCommon.NewCustomError(
		http.StatusBadRequest,
		errors.New("sum of ratio is not 100%"),
		"sum of ratio is not 100%",
		"ErrInvalidRatio",
	)
)
