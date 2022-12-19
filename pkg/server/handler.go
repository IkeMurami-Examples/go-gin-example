package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealhckeckHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func PostHandler(ctx *gin.Context) {
	// PostHandler
}
