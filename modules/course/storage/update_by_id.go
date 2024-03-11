package coursestorage

import (
	"SchoolManagement-BE/appCommon"
	coursemodel "SchoolManagement-BE/modules/course/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (s *MgDBStorage) Update(ctx context.Context, updateData *coursemodel.CourseUpdate) error {
	db := s.db.Database(appCommon.MainDBName).Collection(coursemodel.Course{}.TableName())
	objId, err := primitive.ObjectIDFromHex(updateData.CourseId)

	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}

	conditions := bson.M{
		"_id": objId,
	}

	data := bson.M{}

	if updateData.AttendanceRatio != nil {
		data["attendance_ratio"] = *updateData.AttendanceRatio
	}
	if updateData.LabRatio != nil {
		data["lab_ratio"] = *updateData.LabRatio
	}
	if updateData.MidtermRatio != nil {
		data["midterm_ratio"] = *updateData.MidtermRatio
	}
	if updateData.FinalRatio != nil {
		data["final_ratio"] = *updateData.FinalRatio
	}
	if updateData.Limit != nil {
		data["limit"] = *updateData.Limit
	}
	if updateData.CourseName != nil {
		data["course_name"] = *updateData.CourseName
	}
	if updateData.Credit != nil {
		data["credit"] = *updateData.Credit
	}
	if updateData.Description != nil {
		data["description"] = *updateData.Description
	}

	if updateData.StartTime != nil {
		data["start_time"] = time.Unix(*updateData.StartTime, 0)
	}
	if updateData.EndTime != nil {
		data["end_time"] = time.Unix(*updateData.EndTime, 0)
	}

	if updateData.Period != nil {
		data["period"] = *updateData.Period
	}

	data["updated_at"] = time.Now()

	_, err = db.UpdateOne(ctx, conditions, bson.M{
		"$set": data,
	})

	if err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
