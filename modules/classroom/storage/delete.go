package classroomstorage

import (
	"SchoolManagement-BE/appCommon"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *MgDBStorage) Delete(ctx context.Context, ids []string) error {
	db := s.db.Database(appCommon.MainDBName).Collection(classroommodel.Classroom{}.TableName())

	objIds := make([]primitive.ObjectID, len(ids))
	for i, v := range ids {
		objId, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			return appCommon.ErrInvalidRequest(err)
		}
		objIds[i] = objId
	}

	_, err := db.DeleteMany(ctx, bson.M{
		"_id": bson.M{
			"$in": objIds,
		},
	})
	if err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
