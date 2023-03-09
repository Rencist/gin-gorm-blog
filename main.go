package main

import (
	"gin-gorm-blog/common"
	"gin-gorm-blog/config"
	"gin-gorm-blog/controller"
	"gin-gorm-blog/repository"
	"gin-gorm-blog/routes"
	"gin-gorm-blog/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		res := common.BuildErrorResponse("Gagal Terhubung ke Server", err.Error(), common.EmptyObj{})
		(*gin.Context).JSON((&gin.Context{}), http.StatusBadGateway, res)
		return
	}

	var (
		db *gorm.DB = config.SetupDatabaseConnection()
		
		jwtService service.JWTService = service.NewJWTService()

		userRepository repository.UserRepository = repository.NewUserRepository(db)
		userService service.UserService = service.NewUserService(userRepository)
		userController controller.UserController = controller.NewUserController(userService, jwtService)

		commentRepository repository.CommentRepository = repository.NewCommentRepository(db)
		
		blogRepository repository.BlogRepository = repository.NewBlogRepository(db, commentRepository)
		blogService service.BlogService = service.NewBlogService(blogRepository)
		blogController controller.BlogController = controller.NewBlogController(blogService, jwtService)

		commentService service.CommentService = service.NewCommentService(commentRepository, blogRepository)
		commentController controller.CommentController = controller.NewCommentController(commentService, jwtService)
	)

	server := gin.Default()
	routes.UserRoutes(server, userController, jwtService)
	routes.BlogRoutes(server, blogController, jwtService)
	routes.CommentRoutes(server, commentController, jwtService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	server.Run("127.0.0.1:" + port)
}