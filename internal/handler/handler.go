package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"layout/internal/middleware"
	"layout/pkg/helper/validate"
	"layout/pkg/log"
)

type Handler struct {
	logger *log.Logger
}

func NewHandler(logger *log.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}
func GetUserIdFromCtx(ctx *gin.Context) string {
	v, exists := ctx.Get("claims")
	if !exists {
		return ""
	}
	return v.(*middleware.MyCustomClaims).UserId
}

func (base *Base) GetUserId(c *gin.Context) uint64 {
	if uid, ok := c.Get("u_id"); ok {
		return uid.(uint64)
	}
	return 0
}

func (base *Base) GetUserInfo(c *gin.Context) dto.LoginUserInfo {
	uInfo, _ := c.Get("u_info")
	return uInfo.(dto.LoginUserInfo)
}

// ShouldBind 绑定并翻译错误信息 翻译错误信息
// todo  利用panic这种方法减少 nil ！= error 还待考量 符合语言特性不
func (base *Base) ShouldBind(c *gin.Context, obj any) {
	err := c.ShouldBind(obj)
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, err := range errs {
			panic(validate.NewValidateError(err.Translate(validate.Trans)))
		}
	}
	if err != nil {
		panic(validate.NewValidateError("参数错误"))
	}
}
func (base *Base) GetPageParams(c *gin.Context) (int, int) {
	page := cast.ToInt(c.Query("page"))
	if page <= 0 {
		page = 1
	}
	pageSize := cast.ToInt(c.Query("page_size"))
	if pageSize <= 0 {
		pageSize = 16
	}
	return page, pageSize
}
