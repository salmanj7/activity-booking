package route

import (
    "github.com/gin-gonic/gin"
    "github.com/salmanj7/activity-booking/controller"
	"github.com/salmanj7/activity-booking/middleware"
)

func RegisterRoutes(r *gin.Engine) {
    userRoutes := r.Group("/users")
    {
        userRoutes.GET("/", controller.GetAllUsers)
        userRoutes.POST("/", controller.CreateUser)
    }

    r.POST("/login", controller.Login)

    // Protected routes
    auth := r.Group("/api")
    auth.Use(middleware.AuthMiddleware())
    {
        auth.GET("/profile", controller.GetProfile) // example protected route
    }
}
