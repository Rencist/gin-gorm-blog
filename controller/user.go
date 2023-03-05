package controller

import (
	"gin-gorm-blog/common"
	"gin-gorm-blog/dto"
	"gin-gorm-blog/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	RegisterUser(ctx *gin.Context)
	GetAllUser(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(us service.UserService) UserController {
	return &userController{
		userService: us,
	}
}

func(uc *userController) RegisterUser(ctx *gin.Context) {
	var user dto.UserCreateDto
	err := ctx.ShouldBind(&user)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan User", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, err := uc.userService.RegisterUser(ctx.Request.Context(), user)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan User", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Menambahkan User", result)
	ctx.JSON(http.StatusOK, res)
}

func(uc *userController) GetAllUser(ctx *gin.Context) {
	result, err := uc.userService.GetAllUser(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List User", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan List User", result)
	ctx.JSON(http.StatusOK, res)
}