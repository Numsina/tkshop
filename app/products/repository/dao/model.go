package dao

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
)

type ImageList []string

func (i *ImageList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), i)
}

func (i ImageList) Value() (driver.Value, error) {
	return json.Marshal(i)
}

type Categorys struct {
	Id             int32      `gorm:"column:id;primaryKey;autoIncrement;not null"`
	Name           string     `gorm:"column:name;type:varchar(30);not null"`
	Level          int        `gorm:"column:level;type:int;index;not null;default:1;comment:'1代表一级分类, 2代表二级分类, 3代表三级分类'"`
	ParentId       int32      `gorm:"column:parentId;default:0;not null"`
	ParentCategory *Categorys `gorm:"ForeignKey:ParentId;AssociationForeignKey:Id;constraint:OnDelete:CASCADE"`
	RootId         int32      `gorm:"column:root_id;default:0;index;not null"`
	CreateAt       int64      `gorm:"column:create_at"`
	UpdateAt       int64      `gorm:"column:update_at"`
	DeleteAt       int64      `gorm:"column:delete_at"`
}

type BrandCategorys struct {
	Id         int32 `gorm:"column:id;primaryKey;autoIncrement;not null"`
	BrandId    int32 `gorm:"column:brand_id;uniqueIndex:idx_brand_category;not null"`
	CategoryId int32 `gorm:"column:category_id;uniqueIndex:idx_brand_category;not null"`
	CreateAt   int64 `gorm:"column:create_at"`
	UpdateAt   int64 `gorm:"column:update_at"`
	DeleteAt   int64 `gorm:"column:delete_at"`
}

type Brands struct {
	Id       int32  `gorm:"column:id;primaryKey;autoIncrement;not null"`
	Name     string `gorm:"column:name;type:varchar(30);unique;not null"`
	Logo     string `gorm:"column:logo;type:varchar(200);not null"`
	CreateAt int64  `gorm:"column:create_at"`
	UpdateAt int64  `gorm:"column:update_at"`
	DeleteAt int64  `gorm:"column:delete_at"`
}

type Products struct {
	Id          int32     `gorm:"column:id;primaryKey;autoIncrement;not null"`
	Name        string    `gorm:"column:name;type:varchar(30);not null"`
	CategoryId  int32     `gorm:"column:category_id;not null"`
	BrandId     int32     `gorm:"column:brand_id;not null"`
	Description string    `gorm:"column:description;type:varchar(255);not null"`
	IsNew       bool      `gorm:"column:is_new;not null"`
	IsHot       bool      `gorm:"column:is_hot;not null"`
	OnSale      bool      `gorm:"column:on_sale;not null"`
	Click       int32     `gorm:"column:click;default:0;not null"`
	Sale        int32     `gorm:"column:sale;default:0;not null"`
	Favorite    int32     `gorm:"column:favorite;default:0;not null"`
	MarkPrice   float32   `gorm:"column:mark_price;not null"`
	ShopPrice   float32   `gorm:"column:shop_price;not null"`
	Picture     string    `gorm:"column:picture;not null"`
	Images      ImageList `gorm:"column:images;not null"`
	CreateAt    int64     `gorm:"column:create_at"`
	UpdateAt    int64     `gorm:"column:update_at"`
	DeleteAt    int64     `gorm:"column:delete_at"`
}

func InitTable(db *gorm.DB) {
	db.AutoMigrate(&Categorys{}, &Brands{}, &BrandCategorys{}, &Products{})
}
