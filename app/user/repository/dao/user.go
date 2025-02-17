package dao

import (
	"context"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type UserI interface {
	CreateUser(ctx context.Context, user User) (int32, error)
	UpdateUserInfoByUid(ctx context.Context, user User) (User, error)
	FindUserByEmail(ctx context.Context, email string) (User, error)
}

var _ UserI = &user{}

type user struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserI {
	return &user{db: db}
}

func (u *user) CreateUser(ctx context.Context, user User) (int32, error) {
	now := time.Now().UnixMilli()
	user.CreateAt = now
	user.UpdateAt = now
	err := u.db.WithContext(ctx).Create(&user).Error
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		const uniqueConflictErr uint16 = 1062
		if mysqlErr.Number == uniqueConflictErr {
			return 0, nil
		}
	}

	if err != nil {
		// 可能是数据库错误， 记录日志，
		// todo
	}
	return user.Id, nil
}

func (u *user) UpdateUserInfoByUid(ctx context.Context, user User) (User, error) {
	now := time.Now().UnixMilli()
	user.CreateAt = now
	user.UpdateAt = now
	err := u.db.WithContext(ctx).Updates(&user).Error
	if err != nil {
		// 可能是数据库错误， 记录日志，
		// todo
	}

	return user, nil
}

func (u *user) FindUserByEmail(ctx context.Context, email string) (User, error) {
	var ue = User{}
	if err := u.db.WithContext(ctx).First(&ue, email).Error; err != nil {
		// 记录日志
		return ue, err
	}

	return ue, nil
}
