package classroomstorage

import (
	"SchoolManagement-BE/appCommon"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *MgDBStorage) UpdateStudentScore(
	ctx context.Context,
	data *classroommodel.ClassroomStudentScoreUpdate,
) error {
	db := s.db.Database(appCommon.MainDBName).Collection(classroommodel.Classroom{}.TableName())
	classID, err := primitive.ObjectIDFromHex(data.ClassroomID)
	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}
	userID, err := primitive.ObjectIDFromHex(data.UserID)
	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}
	condition := bson.M{
		"_id":             classID,
		"members.user_id": userID,
	}

	update := bson.M{}

	if data.Attendance != nil {
		update["members.$.attendance"] = *data.Attendance
	}
	if data.Lab != nil {
		update["members.$.lab"] = *data.Lab
	}
	if data.Midterm != nil {
		update["members.$.midterm"] = *data.Midterm
	}
	if data.Final != nil {
		update["members.$.final"] = *data.Final
	}

	if _, err := db.UpdateOne(ctx, condition, bson.M{"$set": update}); err != nil {
		return appCommon.ErrDB(err)
	}

	return nil
}
