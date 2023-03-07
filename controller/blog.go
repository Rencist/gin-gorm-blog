package controller

import (
	"gin-gorm-blog/common"
	"gin-gorm-blog/dto"
	"gin-gorm-blog/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BlogController interface {
	CreateBlog(ctx *gin.Context)
	GetAllBlog(ctx *gin.Context)
	GetUserBlog(ctx *gin.Context)
	GetBlogByID(ctx *gin.Context)
	LikeBlogByID(ctx *gin.Context)
	UpdateBlog(ctx *gin.Context)
}

type blogController struct {
	jwtService service.JWTService
	blogService service.BlogService
}

func NewBlogController(bs service.BlogService, jwts service.JWTService) BlogController {
	return &blogController{
		blogService: bs,
		jwtService: jwts,
	}
}

func(bc *blogController) CreateBlog(ctx *gin.Context) {
	var blog dto.BlogCreateDto
	err := ctx.ShouldBind(&blog)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan Blog", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	token := ctx.MustGet("token").(string)
	userID, err := bc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	blog.UserID = userID
	result, err := bc.blogService.CreateBlog(ctx.Request.Context(), blog)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan Blog", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Menambahkan Blog", result)
	ctx.JSON(http.StatusOK, res)
}

func(bc *blogController) GetAllBlog(ctx *gin.Context) {
	result, err := bc.blogService.GetAllBlog(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Blog", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan List Blog", result)
	ctx.JSON(http.StatusOK, res)
}

func(bc *blogController) GetUserBlog(ctx *gin.Context) {
	UserID := ctx.Param("id")
	result, err := bc.blogService.GetUserBlog(ctx.Request.Context(), UserID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Blog", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan List Blog", result)
	ctx.JSON(http.StatusOK, res)
}

func(bc *blogController) GetBlogByID(ctx *gin.Context) {
	BlogID := ctx.Param("id")
	result, err := bc.blogService.GetBlogByID(ctx.Request.Context(), BlogID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan Blog", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	if result.Title == "" {
		res := common.BuildErrorResponse("Gagal Mendapatkan Blog", "Blog Tidak Ditemukan", common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan Blog", result)
	ctx.JSON(http.StatusOK, res)
}

func(bc *blogController) LikeBlogByID(ctx *gin.Context) {
	BlogID := ctx.Param("id")
	err := bc.blogService.LikeBlogByID(ctx.Request.Context(), BlogID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Like Blog", "Blog Tidak Ditemukan", common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Like Blog", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func(bc *blogController) UpdateBlog(ctx *gin.Context) {
	blogID := ctx.Param("id")
	var blog dto.BlogUpdateDto
	err := ctx.ShouldBind(&blog)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Blog", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	blog.ID, _ = uuid.Parse(blogID)
	token := ctx.MustGet("token").(string)
	userID, err := bc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	checkBlogUser := bc.blogService.ValidateBlogUser(ctx, userID.String(), blogID)
	if !checkBlogUser {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Akun Anda Tidak Memiliki Akses Untuk Mengupdate Blog Ini", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	err = bc.blogService.UpdateBlog(ctx.Request.Context(), blog)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Blog", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Mengupdate Blog", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}