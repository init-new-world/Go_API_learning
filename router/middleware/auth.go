package middleware

import (
	"net/http"

	"github.com/init-new-world/Go_API_learning/pkg/token"

	"github.com/lexkong/log"

	"github.com/gin-gonic/gin"
	"github.com/init-new-world/Go_API_learning/pkg/errno"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if _, err := token.ParseRequest(ctx); err != nil {
			log.Errorf(err, "Token invalid.")
			ctx.JSON(http.StatusOK, errno.ErrorJSON(err))
			ctx.Abort()
		}

		ctx.Next()
	}
}
