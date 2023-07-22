package service

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"layout/global"
	"layout/internal/model"
	"layout/internal/repository"
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

func (s *userService) Login(ctx context.Context, req *LoginRequest) (string, error) {
	user, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil || user == nil {
		return "", errors.Wrap(err, "failed to get user by username")
	}
	//err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if req.Password != "123456" {
		return "", errors.Wrap(err, "failed to hash password")
	}
	token, _ := s.GenerateToken(ctx, user)
	return token, nil
}

// GetProfile 获取用户信息
func (s *userService) GetProfile(ctx context.Context, userId uint64) (*model.User, error) {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user by ID")
	}
	return user, nil
}

// UpdateProfile 修改用户信息
func (s *userService) UpdateProfile(ctx context.Context, userId uint64, req *UpdateProfileRequest) error {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return errors.Wrap(err, "failed to get user by ID")
	}
	user.Mail = req.Email
	user.Nickname = req.Nickname
	if err = s.userRepo.Update(ctx, user); err != nil {
		return errors.Wrap(err, "failed to update user")
	}
	return nil
}

// GenerateToken 生成用户token
func (s *userService) GenerateToken(ctx context.Context, userInfo *model.User) (string, error) {
	channel := "h5"            //此处演示写死
	var duration time.Duration //此处演示写死
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
	if oldToken, _ := global.Redis.HGet(context.Background(), channel, strUserId).Result(); oldToken != "" {
		global.Redis.Del(context.Background(), oldToken)
	}
	global.Redis.Set(context.Background(), token, jsonStr, duration*time.Second)
	global.Redis.HSet(context.Background(), channel, strUserId, token)
	return token, nil
}
