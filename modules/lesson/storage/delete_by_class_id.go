package lessonstorage

import (
	"SchoolManagement-BE/appCommon"
	lessonmodel "SchoolManagement-BE/modules/lesson/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *MgDBStorage) DeleteLessonByClassID(ctx context.Context, classID []string) error {
	db := s.db.Database(appCommon.MainDBName).Collection(lessonmodel.Lesson{}.TableName())
	objIds := make([]primitive.ObjectID, len(classID))

	for i, v := range classID {
		objId, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			return appCommon.ErrInvalidRequest(err)
		}
		objIds[i] = objId

	}

	_, err := db.DeleteMany(ctx, bson.M{
		"class_id": bson.M{"$in": objIds},
	})
	if err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
