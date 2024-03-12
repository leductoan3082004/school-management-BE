package lessonmodel

import (
	"SchoolManagement-BE/appCommon"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

const EntityName = "Lesson"

var AllowedExt = []string{
	".pdf",
	".doc",
	".docx",
	".ppt",
	".pptx",
	".xls",
	".xlsx",
	".zip",
	".rar",
	".7z",
	".mp4",
	".avi",
	".mkv",
	".mp3",
	".wav",
	".flac",
	".jpg",
	".jpeg",
	".png",
	".gif",
}

type Material struct {
	LessonID primitive.ObjectID `json:"lesson_id"`
	Key      string             `json:"key" bson:"key"`
	Name     string             `json:"name" bson:"name"`
}

func (Material) TableName() string {
	return "lesson"
}

type Lesson struct {
	appCommon.MgDBModel `bson:",inline"`
	ClassID             primitive.ObjectID `json:"class_id" bson:"class_id"`
	Name                string             `json:"name" bson:"name"`
	Content             string             `json:"content" bson:"content"`
	Materials           []Material         `json:"materials" bson:"materials"`
}

func (Lesson) TableName() string {
	return "lesson"
}

type LessonCreate struct {
	ClassID string `json:"class_id" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type LessonUpdate struct {
	LessonID string  `json:"lesson_id" binding:"required"`
	Name     *string `json:"name"`
	Content  *string `json:"content"`
}

type LessonDelete struct {
	LessonID string `json:"lesson_id" binding:"required"`
}

type LessonList struct {
	ClassID string  `json:"class_id"`
	Query   *string `json:"query"`
}

var (
	ErrMaterialInvalidFormat = appCommon.NewCustomError(
		http.StatusBadRequest,
		errors.New("username has already existed"),
		"username has already existed",
		"ErrMaterialInvalidFormat",
	)
)
