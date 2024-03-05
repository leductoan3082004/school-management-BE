package userstorage

import (
	"SchoolManagement-BE/appCommon"
	usermodel "SchoolManagement-BE/modules/user/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (s *MgDBStorage) UpdatePassword(ctx context.Context, id, password string) error {
	db := s.db.Database(appCommon.MainDBName).Collection(usermodel.User{}.TableName())
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}

	conditions := bson.M{
		"_id": objId,
	}

	data := bson.M{}

	data["password"] = password
	data["updated_at"] = time.Now()

	_, err = db.UpdateOne(ctx, conditions, bson.M{
		"$set": data,
	})

	if err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
