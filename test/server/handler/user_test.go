package handler

import (
	"bytes"
	"io"
	"layout/cmd/server/wireinject"
	"layout/pkg/configParse"
	"layout/pkg/redis"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

//var (
//	userId = os.Setenv("APP_CONF", "../../../config/local.yml")
//	//token  = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiJ5aHM2SGVzZmdGIiwiZXhwIjoxNjkzOTE0ODgwLCJuYmYiOjE2ODYxMzg4ODAsImlhdCI6MTY4NjEzODg4MH0.NnFrZFgc_333a9PXqaoongmIDksNvQoHzgM_IhJM4MQ"
//)

//var hdl *handler.Handler

//func TestMain(m *testing.M) {
//	fmt.Println("begin")
//	os.Setenv("APP_CONF", "../../../config/local.yml")
//	//app, _, err := main.NewApp()
//	//if err != nil {
//	//	panic(err)
//	//}
//	//http.Run(app, fmt.Sprintf(":%d", global.Config.Http.Port))
//	code := m.Run()
//	fmt.Println("test end")
//	os.Exit(code)
//}

func TestUserHandler_Register(t *testing.T) {
	// 关键点1， 使用gin的Router
	r := setupRouter()
	// 关键点2 构造请求body
	data := url.Values{"a": {"1"}, "b": {"2"}}
	reqbody := strings.NewReader(data.Encode())
	req, err := http.NewRequest(http.MethodGet, "/version", reqbody)
	if err != nil {
		t.Fatalf("构建请求失败, err: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// 构造一个记录
	rec := httptest.NewRecorder()
	//关键点4， 调用web服务的方法
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

	//ctrl := gomock.NewController(t)
	//defer ctrl.Finish()
	//params := service.RegisterRequest{
	//	Username: "xxx",
	//	Password: "123456",
	//	Email:    "xxx@gmail.com",
	//}
	//mockUserService := mock_service.NewMockUserService(ctrl)
	//mockUserService.EXPECT().Register(gomock.Any(), &params).Return(nil)
	//router := setupRouter(mockUserService)
	//paramsJson, _ := json.Marshal(params)
	//resp := performRequest(router, "POST", "/register", bytes.NewBuffer(paramsJson))
	//assert.Equal(t, resp.Code, http.StatusOK)
	// Add assertions for the response body if needed
}

//
//func TestUserHandler_Login(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	params := service.LoginRequest{
//		Username: "xxx",
//		Password: "123456",
//	}
//
//	mockUserService := mock_service.NewMockUserService(ctrl)
//	mockUserService.EXPECT().Login(gomock.Any(), &params).Return(token, nil)
//
//	router := setupRouter(mockUserService)
//	paramsJson, _ := json.Marshal(params)
//
//	resp := performRequest(router, "POST", "/login", bytes.NewBuffer(paramsJson))
//
//	assert.Equal(t, resp.Code, http.StatusOK)
//	// Add assertions for the response body if needed
//}
//
//func TestUserHandler_GetProfile(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockUserService := mock_service.NewMockUserService(ctrl)
//	mockUserService.EXPECT().GetProfile(gomock.Any(), userId).Return(&model.User{
//		Nickname: "xxxxx",
//		Mail:     "xxxxx@gmail.com",
//	}, nil)
//
//	router := setupRouter(mockUserService)
//	req, _ := http.NewRequest("GET", "/user", nil)
//	req.Header.Set("Authorization", "Bearer "+token)
//	resp := httptest.NewRecorder()
//
//	router.ServeHTTP(resp, req)
//
//	assert.Equal(t, resp.Code, http.StatusOK)
//	// Add assertions for the response body if needed
//}
//
//func TestUserHandler_UpdateProfile(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	params := service.UpdateProfileRequest{
//		Nickname: "alan",
//		Email:    "alan@gmail.com",
//		Avatar:   "xxx",
//	}
//
//	mockUserService := mock_service.NewMockUserService(ctrl)
//	mockUserService.EXPECT().UpdateProfile(gomock.Any(), userId, &params).Return(nil)
//
//	router := setupRouter(mockUserService)
//	paramsJson, _ := json.Marshal(params)
//
//	req, _ := http.NewRequest("PUT", "/user", bytes.NewBuffer(paramsJson))
//	req.Header.Set("Authorization", "Bearer "+token)
//	req.Header.Set("Content-Type", "application/json")
//	resp := httptest.NewRecorder()
//
//	router.ServeHTTP(resp, req)
//
//	assert.Equal(t, resp.Code, http.StatusOK)
//	// Add assertions for the response body if needed
//}

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
