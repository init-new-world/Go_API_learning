package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func RequestId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Check for incoming header, use it if exists
		requestId := ctx.Request.Header.Get("X-Request-Id")

		// Create request id with UUID4
		if requestId == "" {
			u4 := uuid.NewV4()
			requestId = u4.String()
		}

		// Expose it for use in the application
		ctx.Set("X-Request-Id", requestId)

		// Set X-Request-Id header
		ctx.Writer.Header().Set("X-Request-Id", requestId)
		ctx.Next()
	}
}
