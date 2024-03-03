package cmd

import (
	"SchoolManagement-BE/appCommon"
	"SchoolManagement-BE/cmd/handler"
	"SchoolManagement-BE/plugin/appredis"
	jwtProvider "SchoolManagement-BE/plugin/tokenprovider/jwt"
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"github.com/lequocbinh04/go-sdk/plugin/aws"
	"github.com/lequocbinh04/go-sdk/plugin/storage/sdkmgo"
	"github.com/spf13/cobra"
)

func newService() goservice.Service {

	service := goservice.New(
		goservice.WithName("mindzone"),
		goservice.WithVersion("1.0.0"),
		goservice.WithInitRunnable(sdkmgo.NewMongoDB("mongodb", appCommon.DBMain)),
		goservice.WithInitRunnable(appredis.NewRedisDB("redis", appCommon.PluginRedis)),
		goservice.WithInitRunnable(jwtProvider.NewJwtProvider("jwt", appCommon.PluginJWT)),
		goservice.WithInitRunnable(aws.New("aws", appCommon.PluginAWS)),
	)

	if err := service.Init(); err != nil {
		panic(err)
	}
	return service
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start a backend service",
	Run: func(cmd *cobra.Command, args []string) {
		service := newService()

		service.HTTPServer().AddHandler(func(engine *gin.Engine) {
			handler.MainRoute(engine, service)
		})

		//go startGRPCService(service)
		if err := service.Start(func() {}); err != nil {
			panic(err)
		}

	},
}

func Execute() {
	rootCmd.AddCommand(outEnvCmd)
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
