package app

import (
	"github.com/gin-gonic/gin"
	"layout/internal/handler"
	"layout/internal/response"
	"layout/internal/service"
)

type UserHandler interface {
	Login(ctx *gin.Context)
	GetProfile(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
}

type userHandler struct {
	*handler.Handler
	userService service.UserService
}

func NewUserHandler(handler *handler.Handler, userService service.UserService) UserHandler {
	return &userHandler{
		Handler:     handler,
		userService: userService,
	}
}

// @Tags 前台用户信息
// @Summary  用户登录注册
// @Accept    application/json
// @Produce   application/json
// @Param object body service.LoginRequest true "登录注册参数"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /api/user/login [post]
func (h *userHandler) Login(ctx *gin.Context) {
	var req service.LoginRequest
	h.ShouldBind(ctx, &req)
	if token, err := h.userService.Login(ctx, &req); err != nil {
		response.FailWithError(ctx, err)
		return
	} else {
		response.OkWithData(ctx, gin.H{
			"accessToken": token,
		})
	}
}

// @Tags 前台用户信息
// @Summary  用户登录注册
// @Accept    application/json
// @Produce   application/json
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /api/user/info [get]
func (h *userHandler) GetProfile(ctx *gin.Context) {
	userId := h.GetUserId(ctx)
	user, err := h.userService.GetProfile(ctx, userId)
	if err != nil {
		response.FailWithError(ctx, err)
		return
	}
	response.OkWithData(ctx, user)
}

// @Tags 前台用户信息
// @Summary  用户登录注册
// @Accept    application/json
// @Produce   application/json
// @Param object body service.UpdateProfileRequest true "参数"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /api/user/update [post]
func (h *userHandler) UpdateProfile(ctx *gin.Context) {
	userId := h.GetUserId(ctx)
	var req service.UpdateProfileRequest
	h.ShouldBind(ctx, &req)
	if err := h.userService.UpdateProfile(ctx, userId, &req); err != nil {
		response.FailWithError(ctx, err)
		return
	}
	response.Ok(ctx)
}
