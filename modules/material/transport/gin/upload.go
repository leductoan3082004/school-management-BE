package gin

import (
	"SchoolManagement-BE/appCommon"
	materialbiz "SchoolManagement-BE/modules/material/biz"
	materialstorage "SchoolManagement-BE/modules/material/storage"
	"errors"
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"github.com/lequocbinh04/go-sdk/plugin/aws"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func UploadByFile(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		lessionID := c.Param("lesson_id")
		fileHeader, err := c.FormFile("file")
		if err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}

		file, err := fileHeader.Open()
		if err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}

		defer file.Close()

		if fileHeader.Size > int64(1024*1024*15) {
			panic(appCommon.ErrInvalidRequest(errors.New("file size too large")))
		}
		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}

		db := sc.MustGet(appCommon.DBMain).(*mongo.Client)
		s3 := sc.MustGet(appCommon.PluginAWS).(aws.S3)

		materialStore := materialstorage.NewMgDBStorage(db)
		materialBiz := materialbiz.NewMaterialUploadBiz(materialStore, s3)

		if err := materialBiz.Upload(c.Request.Context(), dataBytes, fileHeader.Filename, lessionID); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
