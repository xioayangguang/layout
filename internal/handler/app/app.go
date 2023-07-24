package app

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewUserHandler,
)

var StructProvider = wire.Struct(new(Router), "*")

// Router 注册控制器
type Router struct {
	AppUser UserHandler
}
