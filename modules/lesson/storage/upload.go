package lessonstorage

import (
	"SchoolManagement-BE/appCommon"
	lessonmodel "SchoolManagement-BE/modules/lesson/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *MgDBStorage) Upload(ctx context.Context, data *lessonmodel.Material) error {
	db := s.db.Database(appCommon.MainDBName).Collection(lessonmodel.Material{}.TableName())

	condition := bson.M{
		"_id": data.LessonID,
	}

	updateData := bson.M{
		"$push": bson.M{
			"materials": data,
		},
	}

	_, err := db.UpdateOne(ctx, condition, updateData)
	if err != nil {
		return appCommon.ErrDB(err)
	}

	return nil
}
