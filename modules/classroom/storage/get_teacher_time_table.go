package classroomstorage

import (
	"SchoolManagement-BE/appCommon"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *MgDBStorage) GetTeacherTimeTable(ctx context.Context, teacherID string) ([]classroommodel.TimeTable, error) {

	objId, err := primitive.ObjectIDFromHex(teacherID)
	if err != nil {
		return nil, appCommon.ErrInvalidRequest(err)
	}

	db := s.db.Database(appCommon.MainDBName).Collection(classroommodel.Classroom{}.TableName())
	cursor, err := db.Find(ctx, bson.M{
		"teacher_id": objId,
	})
	if err != nil {
		return nil, appCommon.ErrDB(err)
	}
	var result []classroommodel.Classroom
	if err = cursor.All(ctx, &result); err != nil {
		return nil, appCommon.ErrDB(err)
	}

	var timeTable []classroommodel.TimeTable

	fmt.Println(result)
	for i := range result {
		timeTable = append(timeTable, result[i].TimeTable...)
	}
	return timeTable, nil
}
