package coursestorage

import (
	"SchoolManagement-BE/appCommon"
	coursemodel "SchoolManagement-BE/modules/course/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *MgDBStorage) FindById(ctx context.Context, id string) (*coursemodel.Course, error) {
	db := s.db.Database(appCommon.MainDBName).Collection(coursemodel.Course{}.TableName())
	var data coursemodel.Course

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, appCommon.ErrDB(err)
	}
	if err := db.FindOne(ctx, bson.M{"_id": objectId}).Decode(&data); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, appCommon.ErrRecordNotFound
		}
		return nil, appCommon.ErrDB(err)
	}
	return &data, nil
}
