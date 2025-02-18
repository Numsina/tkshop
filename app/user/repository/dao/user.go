package dao

import (
	"context"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserI interface {
	CreateUser(ctx context.Context, user User) (int32, error)
	UpdateUserInfoByUid(ctx context.Context, user User) (User, error)
	FindUserByEmail(ctx context.Context, email string) (User, error)
	DeleteUser(ctx context.Context, uid int32) error
}

var _ UserI = &user{}

type user struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewUserDao(db *gorm.DB, logger *zap.Logger) UserI {
	return &user{
		db:     db,
		logger: logger,
	}
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
			u.logger.Sugar().Infof("唯一主键冲突, 冲突主键: %s", user.Email)
			return 0, nil
		}
	}

	if err != nil {
		// 可能是数据库错误， 记录日志，
		// 这里可以判断一下是否是记录已经存在造成的原因
		u.logger.Sugar().Warnf("记录已存在或数据库错误, 错误原因: %s", err)
	}
	return user.Id, nil
}

func (u *user) DeleteUser(ctx context.Context, uid int32) error {
	return u.db.WithContext(ctx).Delete(&User{Id: uid}).Error
}
func (u *user) UpdateUserInfoByUid(ctx context.Context, user User) (User, error) {
	now := time.Now().UnixMilli()
	user.CreateAt = now
	user.UpdateAt = now
	err := u.db.WithContext(ctx).Updates(&user).Error
	if err != nil {
		// 可能是数据库错误， 记录日志，
		// 这里可以判断一下是否是记录已经存在造成的原因
		u.logger.Sugar().Warnf("记录已存在或数据库错误, 错误原因: %s", err)
	}

	return user, nil
}

func (u *user) FindUserByEmail(ctx context.Context, email string) (User, error) {
	var ue User
	if err := u.db.WithContext(ctx).Where("email = ?", email).First(&ue).Error; err != nil {
		// 记录日志
		u.logger.Sugar().Warnf("记录未找到或者数据库内部错误, 错误原因：%s", err)
		return ue, err
	}

	return ue, nil
}
