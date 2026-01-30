package routes

import (
	"github.com/Raghunandan-79/auth-service/controllers"
	"github.com/Raghunandan-79/auth-service/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	auth := r.Group("/auth") 
	{
		auth.POST("/api/v1/register", controllers.Register)
		auth.POST("/api/v1/login", controllers.Login)
		auth.POST("/api/v1/refresh", controllers.Refresh)
		auth.POST("/api/v1/logout", controllers.Logout)
		auth.GET("/api/v1/me", middleware.Auth(), controllers.Me)
	}

}