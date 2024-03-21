package lessongin

import (
	"SchoolManagement-BE/appCommon"
	"SchoolManagement-BE/component/aws"
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"net/http"
)

func GetMaterial(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Param("key")
		key = key[1:]
		name := c.MustGet(appCommon.CurrentLessonName).(string)
		s3 := sc.MustGet(appCommon.PluginAWS).(aws.S3)
		data, err := s3.GetObjectByKey(c.Request.Context(), key)

		if err != nil {
			panic(appCommon.ErrInternal(err))
		}

		c.Header("Content-Disposition", "attachment; filename="+name)
		c.Data(http.StatusOK, "application/octet-stream", data)

	}
}
