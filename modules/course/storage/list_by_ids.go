package coursestorage

import (
	"SchoolManagement-BE/appCommon"
	coursemodel "SchoolManagement-BE/modules/course/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *MgDBStorage) ListCourseByIDs(
	ctx context.Context,
	CourseIds []string,
) ([]coursemodel.Course, error) {
	collection := s.db.Database(appCommon.MainDBName).Collection(coursemodel.Course{}.TableName())

	courseIds := make([]primitive.ObjectID, len(CourseIds))
	for i, id := range CourseIds {
		oid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return []coursemodel.Course{}, appCommon.ErrInvalidRequest(err)
		}
		courseIds[i] = oid
	}

	cursor, err := collection.Find(ctx, bson.M{
		"_id": bson.M{"$in": courseIds},
	})

	if err != nil {
		return []coursemodel.Course{}, appCommon.ErrDB(err)
	}

	var courses []coursemodel.Course
	if err := cursor.All(ctx, &courses); err != nil {
		return []coursemodel.Course{}, appCommon.ErrDB(err)
	}

	if courses == nil {
		return []coursemodel.Course{}, nil
	}
	return courses, nil
}
