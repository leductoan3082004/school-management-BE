package lessonstorage

import (
	"SchoolManagement-BE/appCommon"
	lessonmodel "SchoolManagement-BE/modules/lesson/model"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *MgDBStorage) DeleteLesson(ctx context.Context, data *lessonmodel.LessonDelete) error {
	db := s.db.Database(appCommon.MainDBName).Collection(lessonmodel.Lesson{}.TableName())
	objId, err := primitive.ObjectIDFromHex(data.LessonID)
	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}

	_, err = db.DeleteOne(ctx, map[string]interface{}{"_id": objId})
	if err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
