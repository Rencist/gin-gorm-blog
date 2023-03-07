package controller

import (
	"gin-gorm-blog/common"
	"gin-gorm-blog/dto"
	"gin-gorm-blog/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CommentController interface {
	CreateComment(ctx *gin.Context)
	UpdateComment(ctx *gin.Context)
}

type commentController struct {
	commentService service.CommentService
	jwtService service.JWTService
}

func NewCommentController(cs service.CommentService, jwts service.JWTService) CommentController {
	return &commentController{
		commentService: cs,
		jwtService: jwts,
	}
}

func(cc *commentController) CreateComment(ctx *gin.Context) {
	blogID := ctx.Param("id")
	var comment dto.CommentCreateDto
	err := ctx.ShouldBind(&comment)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan Comment", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	blogUUID, err := uuid.Parse(blogID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan Comment", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	comment.BlogID = blogUUID
	result, err := cc.commentService.CreateComment(ctx.Request.Context(), comment)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan Comment", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Menambahkan Comment", result)
	ctx.JSON(http.StatusOK, res)
}

func(cc *commentController) UpdateComment(ctx *gin.Context) {
	commentID := ctx.Param("id")
	var comment dto.CommentCreateDto
	err := ctx.ShouldBind(&comment)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Comment", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	
	comment.ID, _ = uuid.Parse(commentID)
	token := ctx.MustGet("token").(string)
	userID, err := cc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	checkCommentUser := cc.commentService.ValidateCommentUser(ctx, userID.String(), commentID)
	if !checkCommentUser {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Akun Anda Tidak Memiliki Akses Untuk Mengupdate Comment Ini", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	err = cc.commentService.UpdateComment(ctx.Request.Context(), comment)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Comment", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Mengupdate Comment", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}