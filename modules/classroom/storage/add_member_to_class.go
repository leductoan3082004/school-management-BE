package classroomstorage

import (
	"SchoolManagement-BE/appCommon"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *MgDBStorage) AddMemberToClass(ctx context.Context, data *classroommodel.Member, classID string) error {
	db := s.db.Database(appCommon.MainDBName).Collection(classroommodel.Classroom{}.TableName())

	classObjID, err := primitive.ObjectIDFromHex(classID)
	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}

	_, err = db.UpdateOne(ctx, bson.M{
		"_id": classObjID,
	}, bson.M{
		"$push": bson.M{
			"members": data,
		},
	})

	if err != nil {
		return appCommon.ErrDB(err)
	}

	return nil
}
