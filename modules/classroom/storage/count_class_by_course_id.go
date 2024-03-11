package classroomstorage

import (
	"SchoolManagement-BE/appCommon"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *MgDBStorage) Count(ctx context.Context, courseID string) (int64, error) {
	objId, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return 0, appCommon.ErrInvalidRequest(err)
	}
	db := s.db.Database(appCommon.MainDBName).Collection(classroommodel.Classroom{}.TableName())
	count, err := db.CountDocuments(ctx, bson.M{
		"course_id": objId,
	})
	if err != nil {
		return 0, appCommon.ErrDB(err)
	}
	return count, nil
}
