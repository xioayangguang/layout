package h5

import "github.com/google/wire"

var ProviderSet = wire.NewSet()

var StructProvider = wire.Struct(new(Router), "*")

// Router 注册控制器
type Router struct {
}
