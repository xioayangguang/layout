## 功能
- **Gin**: https://github.com/gin-gonic/gin
- **Gorm**: https://github.com/go-gorm/gorm
- **Wire**: https://github.com/google/wire
- **Viper**: https://github.com/spf13/viper
- **Zap**: https://github.com/uber-go/zap
- **Golang-jwt**: https://github.com/golang-jwt/jwt
- **Go-redis**: https://github.com/go-redis/redis
- **Testify**: https://github.com/stretchr/testify
- **Sonyflake**: https://github.com/sony/sonyflake
- **gocron**:  https://github.com/go-co-op/gocron

## 目录结构
```
.
├── cmd
│   └── server
│       ├── main.go
│       ├── wire.go
│       └── wire_gen.go
├── config
│   ├── local.yml
│   └── prod.yml
├── internal
│   ├── handler
│   │   ├── handler.go
│   │   └── user.go
│   ├── middleware
│   │   └── cors.go
│   ├── model
│   │   └── user.go
│   ├── repository
│   │   ├── repository.go
│   │   └── user.go
│   ├── server
│   │   └── http.go
│   └── service
│       ├── service.go
│       └── user.go
├── pkg
├── LICENSE
├── README.md
├── README_zh.md
├── go.mod
└── go.sum

```

这是一个经典的Golang 项目的目录结构，包含以下目录：

- cmd: 存放应用程序的入口点，包括主函数和依赖注入的代码。
  - server: 应用程序的主要入口点，包含主函数和依赖注入的代码。
    - main.go: 主函数，用于启动应用程序。
    - wire.go: 使用Wire库生成的依赖注入代码。
    - wire_gen.go: 使用Wire库生成的依赖注入代码。

- config: 存放应用程序的配置文件。
  - local.yml: 本地环境的配置文件。
  - prod.yml: 生产环境的配置文件。

- internal: 存放应用程序的内部代码。
  - handler: 处理HTTP请求的处理程序。
    - handler.go: 处理HTTP请求的通用处理程序。
    - user.go: 处理用户相关的HTTP请求的处理程序。
  - middleware: 存放中间件代码。
    - cors.go: 跨域资源共享中间件。
  - model: 存放数据模型代码。
    - user.go: 用户数据模型。
  - repository: 存放数据访问代码。
    - repository.go: 数据访问的通用接口。
    - user.go: 用户数据访问接口的实现。
  - server: 存放服务器代码。
    - http.go: HTTP服务器的实现。
  - service: 存放业务逻辑代码。
    - service.go: 业务逻辑的通用接口。
    - user.go: 用户业务逻辑的实现。

- pkg: 存放应用程序的公共包。
- storage: 存放应用程序的存储文件。
- go.mod: Go模块文件。
- go.sum: Go模块的依赖版本文件。

