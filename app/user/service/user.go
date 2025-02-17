package service

import (
	"context"
	"github.com/Numsina/tkshop/app/user/domain"
	"github.com/Numsina/tkshop/app/user/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	SignUp(ctx context.Context, user domain.User) (int32, error)
	ModifyUserInfoById(ctx context.Context, user domain.User) (domain.User, error)
	GetUserInfoByEmail(ctx context.Context, email string) (domain.User, error)
}

var _ UserService = &userSvc{}

type userSvc struct {
	userRepo repository.UserRepository
}

func NewUserSvc(userRepo repository.UserRepository) UserService {
	return &userSvc{userRepo: userRepo}
}

func (u *userSvc) SignUp(ctx context.Context, user domain.User) (int32, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		// 记录日志
		return 0, err
	}
	user.Password = string(hash)
	return u.userRepo.CreateUser(ctx, user)
}

func (u *userSvc) ModifyUserInfoById(ctx context.Context, user domain.User) (domain.User, error) {
	if user.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			// 记录日志
			return domain.User{}, err
		}
		user.Password = string(hash)
	}
	return u.userRepo.UpdateUserByUid(ctx, user)
}

func (u *userSvc) GetUserInfoByEmail(ctx context.Context, email string) (domain.User, error) {
	return u.userRepo.GetUserByEmail(ctx, email)
}
