package userstorage

import (
	"SchoolManagement-BE/appCommon"
	usermodel "SchoolManagement-BE/modules/user/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *MgDBStorage) ListDataWithCondition(ctx context.Context, filter *usermodel.UserList, paging *appCommon.Paging, moreInfo ...string) ([]usermodel.User, error) {
	collection := s.db.Database(appCommon.MainDBName).Collection(usermodel.User{}.TableName())

	opts := options.Find()
	if paging == nil {
		paging = &appCommon.Paging{
			Page:       1,
			FakeCursor: "",
			Limit:      50,
		}
	}

	condition := bson.M{}
	if filter.Role != nil {
		condition["role"] = *filter.Role
	}

	// If FakeCursor is given use it for pagination
	if v := paging.FakeCursor; v != "" {
		oid, err := primitive.ObjectIDFromHex(v)
		if err == nil {
			condition["_id"] = bson.M{"$lt": oid}
		}
	} else {
		// Skip the number of documents according to the current page number
		opts.SetSkip(int64((paging.Page - 1) * paging.Limit))
	}

	opts.SetLimit(int64(paging.Limit)).SetSort(bson.D{{"_id", -1}})

	cursor, err := collection.Find(ctx, condition, opts)
	if err != nil {
		return nil, appCommon.ErrDB(err)
	}

	// Get total count
	count, err := collection.CountDocuments(ctx, condition)
	if err != nil {
		return nil, appCommon.ErrDB(err)
	}
	paging.Total = count
	var res []usermodel.User
	if err = cursor.All(ctx, &res); err != nil {
		return nil, appCommon.ErrDB(err)
	}
	return res, nil
}
