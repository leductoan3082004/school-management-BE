package classroommodel

import (
	"SchoolManagement-BE/appCommon"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const EntityName = "Classroom"

type TimeTable struct {
	ID          primitive.ObjectID `json:"id" bson:"id"`
	LessonStart time.Time          `json:"lesson_start" bson:"lesson_start"`
	LessonEnd   time.Time          `json:"lesson_end" bson:"lesson_end"`
}

type Classroom struct {
	appCommon.MgDBModel `json:",inline" bson:",inline"`
	CourseID            primitive.ObjectID `json:"course_id" bson:"course_id"`
	TeacherID           primitive.ObjectID `json:"teacher_id" bson:"teacher_id"`
	TimeTable           []TimeTable        `json:"time_table" bson:"time_table"`
	Limit               int                `json:"limit" bson:"limit"`
}

func (Classroom) TableName() string {
	return "classroom"
}

type ClassroomCreate struct {
	CourseID  string `json:"course_id" binding:"required"`
	TeacherID string `json:"teacher_id" binding:"required"`
	Weeks     int    `json:"weeks" binding:"required"`
	Limit     int    `json:"limit" binding:"required"`
	TimeTable []struct {
		LessonStart int64 `json:"lesson_start" binding:"required"`
		LessonEnd   int64 `json:"lesson_end" binding:"required"`
	} `json:"time_table" binding:"required"`
}

type ClassroomUpdate struct {
	ClassroomId primitive.ObjectID  `json:"classroom_id" binding:"required"`
	TeacherID   *primitive.ObjectID `json:"teacher_id"`
	Limit       *int                `json:"limit"`
	TimeIds     *[]string           `json:"time_ids"`
}

type ClassroomDelete struct {
	ClassroomIds []string `json:"classroom_ids" binding:"required"`
}
