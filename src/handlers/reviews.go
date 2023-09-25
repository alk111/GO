package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootContextEndpointHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Service is up and running",
		"error":   nil,
		"success": true,
	})
}

func HealthCheckEndpointHandler(ctx *gin.Context) {
	ctx.Header("Content-Length", "2")
	ctx.String(http.StatusOK, "OK")
}

func GetReviewsEndpointHandler(ctx *gin.Context) {
	var data []string
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Found",
		"error":   nil,
		"success": true,
		"data":    data,
	})
}
