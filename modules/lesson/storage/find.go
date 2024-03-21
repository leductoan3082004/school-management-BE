package lessonstorage

import (
	"SchoolManagement-BE/appCommon"
	lessonmodel "SchoolManagement-BE/modules/lesson/model"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *MgDBStorage) Find(ctx context.Context, lessonID string) (*lessonmodel.Lesson, error) {
	db := s.db.Database(appCommon.MainDBName).Collection(lessonmodel.Lesson{}.TableName())

	var data lessonmodel.Lesson

	objId, err := primitive.ObjectIDFromHex(lessonID)
	if err != nil {
		return nil, appCommon.ErrInvalidRequest(err)
	}

	err = db.FindOne(ctx, map[string]interface{}{"_id": objId}).Decode(&data)
	if err != nil {
		return nil, appCommon.ErrDB(err)
	}

	return &data, nil
}
