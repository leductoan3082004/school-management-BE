package lessonstorage

import (
	"SchoolManagement-BE/appCommon"
	lessonmodel "SchoolManagement-BE/modules/lesson/model"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *MgDBStorage) FindLessonByKey(ctx context.Context, key string) (*lessonmodel.Lesson, error) {
	db := s.db.Database(appCommon.MainDBName).Collection(lessonmodel.Lesson{}.TableName())

	var data lessonmodel.Lesson

	err := db.FindOne(ctx, map[string]interface{}{"materials.key": key}).Decode(&data)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, appCommon.ErrRecordNotFound
		}
		return nil, appCommon.ErrDB(err)
	}

	return &data, nil
}
