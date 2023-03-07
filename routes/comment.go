package routes

import (
	"gin-gorm-blog/controller"

	"github.com/gin-gonic/gin"
)

func CommentRoutes(router *gin.Engine, CommentController controller.CommentController) {
	userRoutes := router.Group("/api/comment")
	{
		userRoutes.POST("/:id", CommentController.CreateComment)
	}
}