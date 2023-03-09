package routes

import (
	"gin-gorm-blog/controller"
	"gin-gorm-blog/middleware"
	"gin-gorm-blog/service"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, UserController controller.UserController, jwtService service.JWTService) {
	userRoutes := router.Group("/api/user")
	{
		userRoutes.POST("", UserController.RegisterUser)
		userRoutes.GET("", middleware.Authenticate(jwtService, true), UserController.GetAllUser)
		userRoutes.POST("/login", UserController.LoginUser)
		userRoutes.DELETE("/", middleware.Authenticate(jwtService, false), UserController.DeleteUser)
		userRoutes.PUT("/", middleware.Authenticate(jwtService, false), UserController.UpdateUser)
		userRoutes.GET("/me", middleware.Authenticate(jwtService, false), UserController.MeUser)
	}
}