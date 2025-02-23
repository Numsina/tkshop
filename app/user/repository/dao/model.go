package dao

import "gorm.io/gorm"

type User struct {
	Id          int32  `gorm:"primaryKey, autoIncrement"`
	Email       string `gorm:"unique"`
	Password    string
	NickName    string
	Description string
	Avatar      string
	BirthDay    int64
	CreateAt    int64
	UpdateAt    int64
	DeleteAt    int64
}

func InitAutoMigrateTable(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
