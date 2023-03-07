package routes

import (
	"gin-gorm-blog/controller"
	"gin-gorm-blog/middleware"
	"gin-gorm-blog/service"

	"github.com/gin-gonic/gin"
)

func CommentRoutes(router *gin.Engine, CommentController controller.CommentController, jwtService service.JWTService) {
	userRoutes := router.Group("/api/comment")
	{
		userRoutes.POST("/:id", CommentController.CreateComment)
		userRoutes.PUT("/:id", middleware.Authenticate(jwtService), CommentController.UpdateComment)
	}
}