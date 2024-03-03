package userstorage

import (
	"SchoolManagement-BE/appCommon"
	usermodel "SchoolManagement-BE/modules/user/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *MgDBStorage) FindByUsername(ctx context.Context, username string) (*usermodel.User, error) {
	db := s.db.Database(appCommon.MainDBName).Collection(usermodel.User{}.TableName())
	var data usermodel.User
	err := db.FindOne(ctx, bson.M{"username": username}).Decode(&data)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, appCommon.ErrRecordNotFound
		}
		return nil, appCommon.ErrDB(err)
	}
	return &data, nil
}
