package materialbiz

import (
	"SchoolManagement-BE/appCommon"
	materialmodel "SchoolManagement-BE/modules/material/model"
	"context"
	"fmt"
	"github.com/lequocbinh04/go-sdk/logger"
	"github.com/lequocbinh04/go-sdk/plugin/aws"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"path/filepath"
	"slices"
	"time"
)

type materialUploadStore interface {
	Create(ctx context.Context, data *materialmodel.Material) error
}

type materialUploadBiz struct {
	store  materialUploadStore
	s3     aws.S3
	logger logger.Logger
}

func NewMaterialUploadBiz(store materialUploadStore, s3 aws.S3) *materialUploadBiz {
	return &materialUploadBiz{
		store:  store,
		s3:     s3,
		logger: logger.GetCurrent().GetLogger("MaterialUploadBiz"),
	}
}

func (biz *materialUploadBiz) Upload(ctx context.Context, dataByte []byte, fileName, lessonID string) error {
	fileExt := filepath.Ext(fileName) // "img.jpg" => ".jpg"

	if !slices.Contains(materialmodel.AllowedExt, fileExt) {
		return materialmodel.ErrMaterialInvalidFormat
	}

	lessonId, err := primitive.ObjectIDFromHex(lessonID)
	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}

	originName := fileName

	fileName = fmt.Sprintf(
		"%d%s",
		time.Now().Nanosecond(),
		fileExt,
	) // 9129324893248.jpg

	key := appCommon.Join("/", appCommon.S3Path, fileExt)
	_, err = biz.s3.UploadFileData(ctx, dataByte, key)

	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrInternal(err)
	}

	res := &materialmodel.Material{
		Key:      key,
		Name:     originName,
		LessonID: lessonId,
	}

	if err := biz.store.Create(ctx, res); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotCreateEntity(materialmodel.EntityName, err)
	}
	return nil
}
