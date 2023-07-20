package router

import (
	"github.com/gin-gonic/gin"
	"horse/app/controller/common"
	"horse/middleware"
)

// InitApiRouter 不登陆也不验证签名的路由，通常是一些回调路由
func InitApiRouter(Router *gin.Engine) {
	PublicApiGroup := Router.Group("common")
	PublicApiGroup.Use(middleware.RequestLog())
	PublicApiGroup.Use(middleware.SpeedLimit())
	PublicApiGroup.Use(middleware.Recover())
	{
		indexRouter := PublicApiGroup.Group("horses")
		indexRouter.Any("metadata/:token_id", common.NewNftMetaData().MetaData)
	}
	{
		indexRouter := PublicApiGroup.Group("s3")
		indexRouter.GET("upload-url", common.NewAwsS3Info().S3uploadUrl)
		indexRouter.POST("upload", common.NewAwsS3Info().Upload)
	}
}
