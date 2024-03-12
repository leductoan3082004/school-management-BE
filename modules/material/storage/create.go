package materialstorage

import (
	"SchoolManagement-BE/appCommon"
	materialmodel "SchoolManagement-BE/modules/material/model"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (s *MgDBStorage) Create(ctx context.Context, data *materialmodel.Material) error {
	db := s.db.Database(appCommon.MainDBName).Collection(data.TableName())
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	newData, err := db.InsertOne(ctx, data)
	if err != nil {
		return appCommon.ErrDB(err)
	}
	data.Id = newData.InsertedID.(primitive.ObjectID)
	return nil
}
