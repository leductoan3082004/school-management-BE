package materialstorage

import (
	"SchoolManagement-BE/appCommon"
	materialmodel "SchoolManagement-BE/modules/material/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *MgDBStorage) Delete(ctx context.Context, id string) error {
	db := s.db.Database(appCommon.MainDBName).Collection(materialmodel.Material{}.TableName())

	// sao mày ngu v copilot

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}

	_, err = db.DeleteOne(ctx, bson.M{
		"_id": objId,
	})
	if err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
