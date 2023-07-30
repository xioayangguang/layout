package repository

import (
	"context"
	"github.com/pkg/errors"
	"layout/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id uint64) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	GetMaxSerial(ctx context.Context) int
}

type userRepository struct {
	*Repository
}

func NewUserRepository(r *Repository) UserRepository {
	return &userRepository{
		Repository: r,
	}
}
func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	if err := r.getDb(ctx).WithContext(ctx).Create(user).Error; err != nil {
		return errors.Wrap(err, "failed to create user")
	}
	return nil
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	if err := r.getDb(ctx).WithContext(ctx).Save(user).Error; err != nil {
		return errors.Wrap(err, "failed to update user")
	}
	return nil
}

func (r *userRepository) GetByID(ctx context.Context, userId uint64) (*model.User, error) {
	var user model.User
	if err := r.getDb(ctx).WithContext(ctx).Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get user by ID")
	}
	return &user, nil
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := r.getDb(ctx).Where("nickname = ?", username).First(&user).Error
	return &user, err
}

func (r *userRepository) GetMaxSerial(ctx context.Context) int {
	var u model.User
	r.getDb(ctx).Order("serial desc").Take(&u)
	return int(u.Serial)
}
