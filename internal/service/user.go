package service

import (
	"context"
	"encoding/json"
	"layout/global"
	"layout/internal/model"
	"layout/internal/repository"
	"layout/internal/response"
	"layout/pkg/berror"
	"layout/pkg/contextValue"
	"layout/pkg/helper/md5"
	"strconv"
	"time"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateProfileRequest struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email" binding:"required,email"`
	Avatar   string `json:"avatar"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

type UserService interface {
	Login(ctx context.Context, req *LoginRequest) (string, error)
	GetProfile(ctx context.Context, userId uint64) (*model.User, error)
	UpdateProfile(ctx context.Context, userId uint64, req *UpdateProfileRequest) error
	GenerateToken(ctx context.Context, userInfo *model.User) (string, error)
}

type userService struct {
	userRepo repository.UserRepository
	*Service
}

func NewUserService(service *Service, userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
		Service:  service,
	}
}

// Login 登录
func (s *userService) Login(ctx context.Context, req *LoginRequest) (string, error) {
	user, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil || user == nil {
		return "", berror.New(response.LoginError)
	}
	if req.Password == "123456" {
		token, _ := s.GenerateToken(ctx, user)
		return token, nil
	}
	if req.Password == "1234567" {
		panic("你的系统崩溃了")
	}
	return "", berror.New(response.LoginError)
}

// GetProfile 获取用户信息
func (s *userService) GetProfile(ctx context.Context, userId uint64) (*model.User, error) {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return nil, berror.New(response.Error)
	}
	return user, nil
}

// UpdateProfile 修改用户信息
func (s *userService) UpdateProfile(ctx context.Context, userId uint64, req *UpdateProfileRequest) error {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return berror.New(response.Error)
	}
	user.Mail = req.Email
	user.Nickname = req.Nickname
	if err = s.userRepo.Update(ctx, user); err != nil {
		return berror.New(response.Error)
	}
	return nil
}

// GenerateToken 生成用户token
func (s *userService) GenerateToken(ctx context.Context, userInfo *model.User) (string, error) {
	channel := "app"                        //此处演示写死
	var duration time.Duration = 86400 * 30 //此处演示写死
	token := md5.Md5(strconv.Itoa(int(time.Now().UnixNano())) + strconv.Itoa(int(userInfo.Id)))
	jsonStr, _ := json.Marshal(contextValue.LoginUserInfo{
		Id:             userInfo.Id,
		Nickname:       userInfo.Nickname,
		Uuid:           userInfo.Uuid,
		InvitationCode: userInfo.InvitationCode,
		ApiAuth:        token,
		Serial:         userInfo.Serial,
	})
	strUserId := strconv.FormatUint(userInfo.Id, 10)
	if oldToken, _ := global.Redis.HGet(ctx, channel, strUserId).Result(); oldToken != "" {
		global.Redis.Del(ctx, oldToken)
	}
	global.Redis.Set(ctx, token, jsonStr, duration*time.Second)
	global.Redis.HSet(ctx, channel, strUserId, token)
	return token, nil
}
