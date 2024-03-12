package lessonmodel

import (
	"SchoolManagement-BE/appCommon"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const EntityName = "Lesson"

type Lesson struct {
	appCommon.MgDBModel `bson:",inline"`
	ClassID             primitive.ObjectID `json:"class_id" bson:"class_id"`
	CourseID            primitive.ObjectID `json:"course_id" bson:"course_id"`
	Name                string             `json:"name" bson:"name"`
	Content             string             `json:"content" bson:"content"`
}

func (Lesson) TableName() string {
	return "lesson"
}

type LessonCreate struct {
	ClassID  *string `json:"class_id"`
	CourseID *string `json:"course_id"`
	Name     string  `json:"name" binding:"required"`
	Content  string  `json:"content" binding:"required"`
}

type LessonUpdate struct {
	LessonID string  `json:"lesson_id" binding:"required"`
	ClassID  *string `json:"class_id"`
	CourseID *string `json:"course_id"`
	Name     *string `json:"name" binding:"required"`
	Content  *string `json:"content" binding:"required"`
}

type LessonDelete struct {
	LessonID string `json:"lesson_id" binding:"required"`
}

type LessonList struct {
	ClassID  *string `json:"class_id"`
	CourseID *string `json:"course_id"`
}
