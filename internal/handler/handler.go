package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/spf13/cast"
	"layout/internal/validate"
	"layout/pkg/contextValue"
)

var StructProvider = wire.Struct(new(Router), "*")

type Router struct { //注册控制器
	UserAPI UserHandler
}

var ProviderSet = wire.NewSet( //放入容器
	NewHandler,
	NewUserHandler,
)

func NewHandler() *Handler {
	return &Handler{}
}

type Handler struct {
}

func (base *Handler) GetUserId(c *gin.Context) uint64 {
	if uid, ok := c.Get("u_id"); ok {
		return uid.(uint64)
	}
	return 0
}

func (base *Handler) GetUserInfo(c *gin.Context) contextValue.LoginUserInfo {
	uInfo, _ := c.Get("u_info")
	return uInfo.(contextValue.LoginUserInfo)
}

// ShouldBind 绑定并翻译错误信息 翻译错误信息
// todo  利用panic这种方法减少 nil ！= error 还待考量 符合语言特性不
func (base *Handler) ShouldBind(c *gin.Context, obj any) {
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
func (base *Handler) GetPageParams(c *gin.Context) (int, int) {
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
