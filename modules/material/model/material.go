package materialmodel

import (
	"SchoolManagement-BE/appCommon"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

const EntityName = "Material"

type Material struct {
	appCommon.MgDBModel `bson:",inline"`
	LessonID            primitive.ObjectID `json:"lesson_id" bson:"lesson_id"`
	Key                 string             `json:"key" bson:"key"`
	Name                string             `json:"name" bson:"name"`
}

var AllowedExt = []string{".pdf", ".docx", ".pptx", ".xlsx", ".doc", ".xls", ".ppt", ".jpg", ".jpeg", ".png"}

func (Material) TableName() string {
	return "material"
}

var (
	ErrMaterialInvalidFormat = appCommon.NewCustomError(
		http.StatusBadRequest,
		errors.New("ext must be docx, pptx, pdf, jpg, jpeg, png"),
		"ext must be docx, pptx, pdf, jpg, jpeg, png",
		"ErrMaterialInvalidFormat",
	)
)
