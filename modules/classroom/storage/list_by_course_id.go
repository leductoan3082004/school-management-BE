package classroomstorage

import (
	"SchoolManagement-BE/appCommon"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *MgDBStorage) ListClassroom(
	ctx context.Context,
	paging *appCommon.Paging,
	filter *classroommodel.ClassroomList,
) ([]classroommodel.Classroom, error) {
	db := s.db.Database(appCommon.MainDBName).Collection(classroommodel.Classroom{}.TableName())

	courseId, err := primitive.ObjectIDFromHex(filter.CourseID)
	if err != nil {
		return nil, appCommon.ErrInvalidRequest(err)
	}

	opts := options.Find()
	if paging == nil {
		paging = &appCommon.Paging{
			Page:       1,
			FakeCursor: "",
			Limit:      50,
		}
	}
	condition := bson.M{}

	condition["course_id"] = courseId

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

	cursor, err := db.Find(ctx, condition, opts)

	if err != nil {
		return nil, appCommon.ErrDB(err)
	}

	var result []classroommodel.Classroom
	if err = cursor.All(ctx, &result); err != nil {
		return nil, appCommon.ErrDB(err)
	}

	if result == nil {
		return []classroommodel.Classroom{}, nil
	}
	return result, nil
}
