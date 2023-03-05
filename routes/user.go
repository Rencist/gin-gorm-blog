package routes

import (
	"gin-gorm-blog/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, UserController controller.UserController) {
	userRoutes := router.Group("/api/user")
	{
		userRoutes.POST("", UserController.RegisterUser)
	}
}