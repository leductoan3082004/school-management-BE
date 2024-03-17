package userstorage

import (
	"SchoolManagement-BE/appCommon"
	usermodel "SchoolManagement-BE/modules/user/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *MgDBStorage) FindMemberByIds(ctx context.Context, id []string) ([]usermodel.User, error) {
	db := s.db.Database(appCommon.MainDBName).Collection(usermodel.User{}.TableName())
	var data []usermodel.User

	objId := make([]primitive.ObjectID, len(id))

	for i, v := range id {
		id, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			return []usermodel.User{}, appCommon.ErrInvalidRequest(err)
		}
		objId[i] = id
	}

	cursor, err := db.Find(ctx, bson.M{"_id": bson.M{"$in": objId}})
	if err != nil {
		return []usermodel.User{}, appCommon.ErrDB(err)
	}

	if err = cursor.All(ctx, &data); err != nil {
		return []usermodel.User{}, appCommon.ErrDB(err)
	}

	if data == nil {
		return []usermodel.User{}, nil
	}
	return data, nil
}
