package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GraphQLHandler(ctx *gin.Context) {

}

func HealhckeckHandler(ctx *gin.Context) {
	logger.Info("Called /healthcheck")
	ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func PostHandler(ctx *gin.Context) {
	// PostHandler
}
