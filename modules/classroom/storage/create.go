package classroomstorage

import (
	"SchoolManagement-BE/appCommon"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (s *MgDBStorage) Create(ctx context.Context, data *classroommodel.Classroom) error {
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
