package classroomstorage

import (
	"SchoolManagement-BE/appCommon"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *MgDBStorage) CheckMemberInClass(
	ctx context.Context,
	classID string,
	userID string,
	role *int,
) (bool, error) {
	db := s.db.Database(appCommon.MainDBName).Collection(classroommodel.Classroom{}.TableName())

	classObjID, err := primitive.ObjectIDFromHex(classID)
	if err != nil {
		return false, appCommon.ErrInvalidRequest(err)
	}

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return false, appCommon.ErrInvalidRequest(err)
	}

	condition := bson.M{
		"user_id": userObjID,
	}
	if role != nil {
		condition["role"] = *role
	}
	count, err := db.CountDocuments(ctx, bson.M{
		"_id": classObjID,
		"members": bson.M{
			"$elemMatch": condition,
		},
	})

	if err != nil {
		return false, appCommon.ErrDB(err)
	}

	return count > 0, nil
}
