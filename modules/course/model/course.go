package coursemodel

import (
	"SchoolManagement-BE/appCommon"
	"errors"
	"net/http"
	"time"
)

const EntityName = "Course"

type CourseRatio struct {
	AttendanceRatio int `json:"attendance_ratio" bson:"attendance_ratio"`
	LabRatio        int `json:"lab_ratio" bson:"lab_ratio"`
	MidtermRatio    int `json:"midterm_ratio" bson:"midterm_ratio"`
	FinalRatio      int `json:"final_ratio" bson:"final_ratio"`
}
type CourseSpec struct {
	Limit       int    `json:"limit" bson:"limit"`
	CourseName  string `json:"course_name" bson:"course_name"`
	Credit      int    `json:"credit" bson:"credit"`
	Description string `json:"description" bson:"description"`
	Period      uint   `json:"period" bson:"period"`
}
type CourseRegisterTimeline struct {
	StartTime time.Time `json:"start_time" bson:"start_time"`
	EndTime   time.Time `json:"end_time" bson:"end_time"`
}

type Course struct {
	appCommon.MgDBModel    `json:",inline" bson:",inline"`
	CourseSpec             `json:",inline" bson:",inline"`
	CourseRatio            `json:",inline" bson:",inline"`
	CourseRegisterTimeline `json:",inline" bson:",inline"`
}

func (Course) TableName() string {
	return "course"
}

type CourseUpdate struct {
	CourseId        string  `json:"course_id" binding:"required"`
	AttendanceRatio *uint   `json:"attendance_ratio"`
	LabRatio        *uint   `json:"lab_ratio"`
	MidtermRatio    *uint   `json:"midterm_ratio"`
	FinalRatio      *uint   `json:"final_ratio"`
	Limit           *uint   `json:"limit"`
	CourseName      *string `json:"course_name"`
	Credit          *int    `json:"credit"`
	Description     *string `json:"description"`
	StartTime       *int64  `json:"start_time"`
	EndTime         *int64  `json:"end_time"`
	Period          *uint   `json:"period"`
}
type CourseCreate struct {
	AttendanceRatio uint   `json:"attendance_ratio"`
	LabRatio        uint   `json:"lab_ratio"`
	MidtermRatio    uint   `json:"midterm_ratio"`
	FinalRatio      uint   `json:"final_ratio"`
	Limit           uint   `json:"limit"`
	CourseName      string `json:"course_name" binding:"required"`
	Credit          int    `json:"credit"`
	Description     string `json:"description"`
	StartTime       int64  `json:"start_time" bson:"start_time"`
	EndTime         int64  `json:"end_time" bson:"end_time"`
	Period          uint   `json:"period"`
}

type CourseList struct {
	Query   *string `form:"query"`
	EndTime *int64  `form:"end_time"`
}

type CourseDelete struct {
	CourseId string `json:"course_id" binding:"required"`
}

func (s *CourseCreate) Validate() error {
	if s.FinalRatio+s.AttendanceRatio+s.MidtermRatio+s.LabRatio != 100 {
		return ErrInvalidRatio
	}
	if s.StartTime > s.EndTime {
		return ErrInvalidTimeframe
	}
	return nil
}
func (s *Course) Validate() error {
	if s.FinalRatio+s.AttendanceRatio+s.MidtermRatio+s.LabRatio != 100 {
		return ErrInvalidRatio
	}
	if s.StartTime.After(s.EndTime) {
		return ErrInvalidTimeframe
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
	ErrInvalidTimeframe = appCommon.NewCustomError(
		http.StatusBadRequest,
		errors.New("time is not valid"),
		"time is not valid",
		"ErrInvalidTimeframe",
	)
)
