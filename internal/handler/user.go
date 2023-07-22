package handler

import (
	"github.com/gin-gonic/gin"
	"layout/internal/response"
	"layout/internal/service"
)

type UserHandler interface {
	Login(ctx *gin.Context)
	GetProfile(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
}

type userHandler struct {
	*Handler
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (h *userHandler) Login(ctx *gin.Context) {
	var req service.LoginRequest
	h.ShouldBind(ctx, &req)
	token, err := h.userService.Login(ctx, &req)
	if err != nil {
		response.FailWithError(ctx, err)
		return
	}
	response.OkWithData(ctx, gin.H{
		"accessToken": token,
	})
}

func (h *userHandler) GetProfile(ctx *gin.Context) {
	userId := h.GetUserId(ctx)
	user, err := h.userService.GetProfile(ctx, userId)
	if err != nil {
		response.FailWithError(ctx, nil)
		return
	}
	response.OkWithData(ctx, user)
}

func (h *userHandler) UpdateProfile(ctx *gin.Context) {
	userId := h.GetUserId(ctx)
	var req service.UpdateProfileRequest
	h.ShouldBind(ctx, &req)
	if err := h.userService.UpdateProfile(ctx, userId, &req); err != nil {
		response.FailWithError(ctx, err)
		return
	}
	response.OkWithData(ctx, nil)
}
