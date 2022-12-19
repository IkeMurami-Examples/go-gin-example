package server

import (
	"context"
	"net/http"
	"time"

	"github.com/IkeMurami-Examples/go-gin-example/pkg/server/middleware"
	"github.com/IkeMurami-Examples/go-gin-example/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var logger *zap.Logger

func NewServer(ctx context.Context, endpoint string, mux http.Handler) *http.Server {
	// Logger
	logger = utils.LoggerFromContext(ctx)

	// Creates a router without any middleware by default
	router := gin.New()

	// Middlewares

	// Logger middleware
	router.Use(middleware.Logger(ctx))
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// Non authorized
	root := router.Group("/")
	{
		root.GET("/healthcheck", HealhckeckHandler)

		api := router.Group("/api")
		{
			// Authorized
			// authorized := api.Group("/user")
			// authorized.Use(AuthRequired())
			// {
			//
			// }

			v1 := api.Group("/v1")
			{
				// Non authorized
				v1.GET("/hello", func(ctx *gin.Context) {
					ctx.JSON(http.StatusOK, "hello")
				})

				v1.POST("/post", PostHandler)
			}
		}
	}

	err := router.Run(endpoint)
	if err != nil {
		logger.Error("Gin Server is not started", zap.Error(err))
		return nil
	}

	return &http.Server{
		Addr:           endpoint,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
