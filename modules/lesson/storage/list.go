package lessonstorage

import (
	"SchoolManagement-BE/appCommon"
	lessonmodel "SchoolManagement-BE/modules/lesson/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *MgDBStorage) ListLesson(
	ctx context.Context,
	data *lessonmodel.LessonList,
	paging *appCommon.Paging,
) ([]lessonmodel.Lesson, error) {
	collection := s.db.Database(appCommon.MainDBName).Collection(lessonmodel.Lesson{}.TableName())

	opts := options.Find()
	if paging == nil {
		paging = &appCommon.Paging{
			Page:       1,
			FakeCursor: "",
			Limit:      50,
		}
	}

	condition := bson.M{}

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

	if data.Query != nil {
		condition["$text"] = bson.M{
			"$search": *data.Query,
		}
	}
	classID, err := primitive.ObjectIDFromHex(data.ClassroomID)
	if err != nil {
		return []lessonmodel.Lesson{}, appCommon.ErrInvalidRequest(err)
	}
	condition["classroom_id"] = classID

	opts.SetLimit(int64(paging.Limit)).SetSort(bson.D{{"_id", -1}})

	cursor, err := collection.Find(ctx, condition, opts)
	if err != nil {
		return []lessonmodel.Lesson{}, appCommon.ErrDB(err)
	}

	// Get total count
	count, err := collection.CountDocuments(ctx, condition)
	if err != nil {
		return []lessonmodel.Lesson{}, appCommon.ErrDB(err)
	}
	paging.Total = count
	var res []lessonmodel.Lesson
	if err = cursor.All(ctx, &res); err != nil {
		return []lessonmodel.Lesson{}, appCommon.ErrDB(err)
	}

	if res == nil {
		return []lessonmodel.Lesson{}, nil
	}
	return res, nil
}
