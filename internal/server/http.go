package server

import (
	"github.com/gin-gonic/gin"
	"layout/internal/handler"
	"layout/internal/middleware"
	"layout/internal/router"
	"layout/internal/router/api"
	"layout/internal/router/h5"
	"layout/pkg/helper/resp"
	"layout/pkg/log"
)

func NewServerHTTP(
	logger *log.Logger,
	jwt *middleware.JWT,
	userHandler handler.UserHandler,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	api.InitApiRouter(r)
	h5.InitApiRouter(r)
	router.InitApiRouter(r)
	router.InitExtraRouter(r)

	r.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		//middleware.SignMiddleware(log),
	)

	// 无权限路由
	noAuthRouter := r.Group("/").Use(middleware.RequestLogMiddleware(logger))
	{

		noAuthRouter.GET("/", func(ctx *gin.Context) {
			logger.WithContext(ctx).Info("hello")
			resp.HandleSuccess(ctx, map[string]interface{}{
				"say": "Hi Nunu!",
			})
		})

		noAuthRouter.POST("/register", userHandler.Register)
		noAuthRouter.POST("/login", userHandler.Login)
	}

	// 非严格权限路由
	noStrictAuthRouter := r.Group("/").Use(middleware.NoStrictAuth(jwt, logger), middleware.RequestLogMiddleware(logger))
	{
		noStrictAuthRouter.GET("/user", userHandler.GetProfile)
	}

	// 严格权限路由
	strictAuthRouter := r.Group("/").Use(middleware.StrictAuth(jwt, logger), middleware.RequestLogMiddleware(logger))
	{
		strictAuthRouter.PUT("/user", userHandler.UpdateProfile)
	}

	return r
}
