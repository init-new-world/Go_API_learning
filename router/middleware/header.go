package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Options(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodOptions {
		ctx.Next()
	} else {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		ctx.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		ctx.Header("Content-Type", "application/json")
		ctx.AbortWithStatus(http.StatusOK)
	}
}
