package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"layout/cmd/server/wireinject"
	"layout/internal/service"
	"layout/pkg/configParse"
	"layout/pkg/redis"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	fmt.Println("begin")
	os.Setenv("APP_CONF", "../../../config/local.yml")
	code := m.Run()
	fmt.Println("test end")
	os.Exit(code)
}

func TestUserHandler_Login(t *testing.T) {
	r := setupRouter()
	//构造请求body
	var data = service.LoginRequest{
		Username: "zs",
		Password: "123",
	}
	jsonStr, _ := json.Marshal(data)
	reqbody := strings.NewReader(string(jsonStr))
	req, err := http.NewRequest(http.MethodPost, "/api/user/login", reqbody)
	if err != nil {
		t.Fatalf("构建请求失败, err: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	// 构造一个记录
	rec := httptest.NewRecorder()
	// 调用web服务的方法
	r.ServeHTTP(rec, req)
	result := rec.Result()
	if result.StatusCode != 200 {
		t.Fatalf("请求状态码不符合预期")
	}
	body, err := io.ReadAll(result.Body)
	if err != nil {
		t.Fatalf("读取返回内容失败, err:%v", err)
	}
	defer result.Body.Close()
	t.Log(string(body))
	t.Log("用例测试通过")
}

func setupRouter() *gin.Engine {
	os.Setenv("APP_CONF", "../../../config/local.yml")
	configParse.InitConfig()
	redis.InitRedis()
	app, _, err := wireinject.NewApp()
	if err != nil {
		panic(err)
	}
	return app
}

func performRequest(r http.Handler, method, path string, body *bytes.Buffer) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	return resp
}
