package lessonstorage

import (
	"SchoolManagement-BE/appCommon"
	lessonmodel "SchoolManagement-BE/modules/lesson/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *MgDBStorage) Update(ctx context.Context, data *lessonmodel.LessonUpdate) error {
	db := s.db.Database(appCommon.MainDBName).Collection(lessonmodel.Lesson{}.TableName())

	lessonID, err := primitive.ObjectIDFromHex(data.LessonID)

	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}

	setCondition := bson.M{}

	condition := bson.M{
		"_id": lessonID,
	}

	if data.Name != nil {
		setCondition["name"] = *data.Name
	}
	if data.Content != nil {
		setCondition["content"] = *data.Content
	}

	_, err = db.UpdateOne(ctx, condition, bson.M{
		"$set": setCondition,
	})
	if err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
