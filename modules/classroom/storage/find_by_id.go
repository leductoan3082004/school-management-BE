package classroomstorage

import (
	"SchoolManagement-BE/appCommon"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *MgDBStorage) FindById(ctx context.Context, id string) (*classroommodel.Classroom, error) {
	db := s.db.Database(appCommon.MainDBName).Collection(classroommodel.Classroom{}.TableName())
	var data classroommodel.Classroom

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
