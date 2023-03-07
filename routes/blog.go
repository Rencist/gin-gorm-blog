package routes

import (
	"gin-gorm-blog/controller"
	"gin-gorm-blog/middleware"
	"gin-gorm-blog/service"

	"github.com/gin-gonic/gin"
)

func BlogRoutes(router *gin.Engine, BlogController controller.BlogController, jwtService service.JWTService) {
	userRoutes := router.Group("/api/blog")
	{
		userRoutes.POST("", middleware.Authenticate(jwtService), BlogController.CreateBlog)
		userRoutes.GET("", BlogController.GetAllBlog)
	}
}