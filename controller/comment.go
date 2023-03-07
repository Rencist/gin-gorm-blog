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
}

type commentController struct {
	commentService service.CommentService
}

func NewCommentController(cs service.CommentService) CommentController {
	return &commentController{
		commentService: cs,
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