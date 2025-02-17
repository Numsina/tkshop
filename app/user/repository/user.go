package repository

import (
	"context"

	"github.com/Numsina/tkshop/app/user/domain"
	"github.com/Numsina/tkshop/app/user/repository/dao"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user domain.User) (int32, error)
	UpdateUserByUid(ctx context.Context, user domain.User) (domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
}

var _ UserRepository = &userRepo{}

type userRepo struct {
	dao dao.UserI
}

func NewUserRepository(data dao.UserI) UserRepository {
	return &userRepo{dao: data}
}

func (u *userRepo) CreateUser(ctx context.Context, user domain.User) (int32, error) {
	return u.dao.CreateUser(ctx, u.toDao(user))
}

func (u *userRepo) UpdateUserByUid(ctx context.Context, user domain.User) (domain.User, error) {
	ue, err := u.dao.UpdateUserInfoByUid(ctx, u.toDao(user))
	if err == nil {
		return u.toDomain(ue), nil
	}
	return domain.User{}, err
}

func (u *userRepo) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	ue, err := u.dao.FindUserByEmail(ctx, email)
	if err == nil {
		return u.toDomain(ue), nil
	}
	return domain.User{}, err
}

func (u *userRepo) toDao(user domain.User) dao.User {
	return dao.User{
		Id:          user.Id,
		Email:       user.Email,
		Password:    user.Password,
		NickName:    user.NickName,
		Description: user.Description,
		Avatar:      user.Avatar,
		BirthDay:    user.BirthDay,
	}
}

func (u *userRepo) toDomain(user dao.User) domain.User {
	return domain.User{
		Id:          user.Id,
		Email:       user.Email,
		NickName:    user.NickName,
		Description: user.Description,
		Avatar:      user.Avatar,
		BirthDay:    user.BirthDay,
		CreateAt:    user.CreateAt,
	}
}
