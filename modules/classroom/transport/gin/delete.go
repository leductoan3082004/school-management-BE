package classroomgin

import (
	"SchoolManagement-BE/appCommon"
	classroombiz "SchoolManagement-BE/modules/classroom/biz"
	classroommodel "SchoolManagement-BE/modules/classroom/model"
	classroomstorage "SchoolManagement-BE/modules/classroom/storage"
	lessonstorage "SchoolManagement-BE/modules/lesson/storage"
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"net/http"
)

func Delete(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data classroommodel.ClassroomDelete
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.DBMain).(*mongo.Client)

		classStore := classroomstorage.NewMgDBStorage(db)
		lessonStore := lessonstorage.NewMgDBStorage(db)
		biz := classroombiz.NewDeleteClassroomBiz(classStore, lessonStore)

		wc := writeconcern.New(writeconcern.WMajority())
		txnOptions := options.Transaction().SetWriteConcern(wc)

		session, err := db.StartSession()
		if err != nil {
			panic(err)
		}
		defer session.EndSession(c.Request.Context())
		_, err = session.WithTransaction(c.Request.Context(), func(ctx mongo.SessionContext) (interface{}, error) {
			return nil, biz.DeleteClassroom(ctx, &data)
		}, txnOptions)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))

	}
}
