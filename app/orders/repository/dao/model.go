package dao

import "gorm.io/gorm"

type Orders struct {
	Id       int32          `gorm:"type:int;primaryKey;autoIncrement;not null"`
	UserId   int32          `gorm:"type:int;index;not null"`
	OrderSn  string         `gorm:"type:varchar(50);uniqueIndex;not null"`
	PayType  string         `gorm:"type:varchar(20);not null; comment 'alipay(支付宝), wechat(微信)'"`
	Status   int32          `gorm:"type:int;index;not null;  comment '0(未支付), 1(支付成功), 2(支付失败), 3(超时未支付)'"`
	PayTime  int64          `gorm:"type:int;not null"`
	Amount   float32        `gorm:"type:decimal(10,2);not null"`
	Address  string         `gorm:"type:varchar(255);not null"`
	Phone    string         `gorm:"type:varchar(11);not null"`
	CreateAt int64          `gorm:"type:int;not null"`
	UpdateAt int64          `gorm:"type:int;not null"`
	DeleteAt gorm.DeletedAt `gorm:"type:int;not null"`
}

type OrderGoods struct {
	Id       int32          `gorm:"type:int;primaryKey;autoIncrement;not null"`
	OrderSn  string         `gorm:"index;not null"`
	GoodsId  int32          `gorm:"type:int;not null"`
	Price    float32        `gorm:"type:decimal(10,2);not null"`
	Nums     int32          `gorm:"type:int;not null"`
	CreateAt int64          `gorm:"type:int;not null"`
	UpdateAt int64          `gorm:"type:int;not null"`
	DeleteAt gorm.DeletedAt `gorm:"type:int;not null"`
}

func InitOrderTable(db *gorm.DB) error {
	return db.AutoMigrate(&Orders{}, &OrderGoods{})
}
