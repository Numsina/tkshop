package dao

import (
	"gorm.io/gorm"
)

type Carts struct {
	Id       int32 `gorm:"primaryKey;autoIncrement;not null"`
	UserID   int32 `gorm:"type:int;uniqueIndex:idx_user_goods;not null"`
	GoodsID  int32 `gorm:"type:int;uniqueIndex:idx_user_goods;not null"`
	Nums     int32 `gorm:"type:int;not null"`
	Checked  bool  `gorm:"not null"`
	CreateAt int64
	UpdateAt int64
	DeleteAt gorm.DeletedAt
}

func InitCartsTable(db *gorm.DB) error {
	return db.AutoMigrate(&Carts{})
}
