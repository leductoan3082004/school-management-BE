package coursemodel

import "SchoolManagement-BE/appCommon"

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

type CourseCreate struct {
	AttendanceRatio uint   `json:"attendance_ratio"`
	LabRatio        uint   `json:"lab_ratio"`
	MidtermRatio    uint   `json:"midterm_ratio"`
	FinalRatio      uint   `json:"final_ratio"`
	Limit           uint   `json:"limit"`
	CourseName      string `json:"course_name" binding:"required"`
	Credit          int    `json:"credit"`
}
