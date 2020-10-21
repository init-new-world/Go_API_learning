package router

import (
	"net/http"

	"github.com/init-new-world/Go_API_learning/handler/user"

	"github.com/gin-gonic/gin"
	"github.com/init-new-world/Go_API_learning/handler/sd"
	"github.com/init-new-world/Go_API_learning/router/middleware"
)

func Load(server *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	server.Use(gin.Recovery())
	server.Use(middleware.Options)
	server.Use(mw...)

	server.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 Not Found! Maybe incorrect route.")
	})

	monitor := server.Group("/sd")
	{
		monitor.GET("/health", sd.HealthCheck)
		monitor.GET("/monitor", sd.MonitorCheck)
	}

	Api := server.Group("/api")
	{
		V1 := Api.Group("/v1")
		{
			User := V1.Group("/user").Use(middleware.AuthMiddleware())
			{
				User.POST("/", user.Create)
				User.DELETE("/:username", user.Delete)
				User.PUT("/", user.Update)
				User.GET("/", user.List)
				User.GET("/:username", user.Get)
			}
		}
	}

	server.POST("/login", user.Login)

	return server
}
