package lessonstorage

import (
	"SchoolManagement-BE/appCommon"
	lessonmodel "SchoolManagement-BE/modules/lesson/model"
	"context"
)

func (s *MgDBStorage) ListLesson(
	ctx context.Context,
	data *lessonmodel.LessonList,
	paging *appCommon.Paging,
) ([]lessonmodel.Lesson, error) {
	//TODO implement me
	panic("implement me")
}
