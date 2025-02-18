package service

import (
	"context"

	"github.com/Numsina/tkshop/app/user/domain"
	"github.com/Numsina/tkshop/app/user/repository"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	SignUp(ctx context.Context, user domain.User) (int32, error)
	Login(ctx context.Context, user domain.User) (domain.User, error)
	Delele(ctx context.Context, uid int32) error
	ModifyUserInfoById(ctx context.Context, user domain.User) (domain.User, error)
	GetUserInfoByEmail(ctx context.Context, email string) (domain.User, error)
}

var _ UserService = &userSvc{}

type userSvc struct {
	userRepo repository.UserRepository
	logger   *zap.Logger
}

func NewUserSvc(userRepo repository.UserRepository, logger *zap.Logger) UserService {
	return &userSvc{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (u *userSvc) SignUp(ctx context.Context, user domain.User) (int32, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		// 记录日志
		u.logger.Sugar().Infof("bcrypt加密失败, 失败原因：%s", err)
		return 0, err
	}
	user.Password = string(hash)
	return u.userRepo.CreateUser(ctx, user)
}

func (u *userSvc) Login(ctx context.Context, user domain.User) (domain.User, error) {
	ue, err := u.userRepo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return domain.User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(ue.Password), []byte(user.Password))
	return ue, err
}

func (u *userSvc) Delele(ctx context.Context, uid int32) error {
	return u.userRepo.DeleteUser(ctx, uid)
}

func (u *userSvc) ModifyUserInfoById(ctx context.Context, user domain.User) (domain.User, error) {
	if user.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			// 记录日志
			u.logger.Sugar().Infof("修改密码加密失败, 失败原因：%s", err)
			return domain.User{}, err
		}
		user.Password = string(hash)
	}
	return u.userRepo.UpdateUserByUid(ctx, user)
}

func (u *userSvc) GetUserInfoByEmail(ctx context.Context, email string) (domain.User, error) {
	return u.userRepo.GetUserByEmail(ctx, email)
}
