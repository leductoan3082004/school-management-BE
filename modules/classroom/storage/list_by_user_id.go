package classroomstorage

import (
	"SchoolManagement-BE/appCommon"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *MgDBStorage) ListClassroomByUserID(
	ctx context.Context,
	paging *appCommon.Paging,
	userID string,
) ([]classroommodel.Classroom, error) {
	db := s.db.Database(appCommon.MainDBName).Collection(classroommodel.Classroom{}.TableName())

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return []classroommodel.Classroom{}, appCommon.ErrInvalidRequest(err)
	}

	opts := options.Find()
	if paging == nil {
		paging = &appCommon.Paging{
			Page:       1,
			FakeCursor: "",
			Limit:      50,
		}
	}

	condition := bson.M{
		"members.user_id": userObjID,
	}

	cursor, err := db.Find(ctx, condition, opts)
	if err != nil {
		return []classroommodel.Classroom{}, appCommon.ErrDB(err)
	}

	var result []classroommodel.Classroom
	if err = cursor.All(ctx, &result); err != nil {
		return []classroommodel.Classroom{}, appCommon.ErrDB(err)
	}

	if result == nil {
		return []classroommodel.Classroom{}, nil
	}

	return result, nil
}
