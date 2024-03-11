package classroomstorage

import (
	"SchoolManagement-BE/appCommon"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *MgDBStorage) UpdateClass(ctx context.Context, data *classroommodel.ClassroomUpdate) error {
	db := s.db.Database(appCommon.MainDBName).Collection(classroommodel.Classroom{}.TableName())

	classId, err := primitive.ObjectIDFromHex(data.ClassroomId)

	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}

	setCondition := bson.M{}

	condition := bson.M{
		"_id": classId,
	}
	updateData := bson.M{}
	if data.TeacherID != nil {
		teacherId, err := primitive.ObjectIDFromHex(*data.TeacherID)
		if err != nil {
			return appCommon.ErrInvalidRequest(err)
		}
		updateData["teacher_id"] = teacherId
	}
	if data.Limit != nil {
		updateData["limit"] = *data.Limit
	}

	setCondition["$set"] = updateData

	if data.TimeIds != nil {
		timeIds := make([]primitive.ObjectID, len(*data.TimeIds))
		for i, v := range *data.TimeIds {
			id, err := primitive.ObjectIDFromHex(v)
			if err != nil {
				return appCommon.ErrInvalidRequest(err)
			}
			timeIds[i] = id
		}
		setCondition["$pull"] = bson.M{"time_table": bson.M{"id": bson.M{"$in": timeIds}}}
	}

	_, err = db.UpdateOne(ctx, condition, setCondition)
	if err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
