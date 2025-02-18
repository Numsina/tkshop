package dao

import (
	"context"
	"errors"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrRecordNotFound = errors.New("该记录不存在")
)

type Product interface {
	UpsertProduct(ctx context.Context, p Products) (Products, error)
	DeleteProduct(ctx context.Context, id int32) error
	QueryProductInfoById(ctx context.Context, id int32) (Products, error)
}

var _ Product = &product{}

type product struct {
	db     *gorm.DB
	logger *zap.Logger
}

func (this *product) UpsertProduct(ctx context.Context, p Products) (Products, error) {
	now := time.Now().UnixMilli()
	p.CreateAt = now
	p.UpdateAt = now
	result := this.db.WithContext(ctx).Where("id = ?", p.Id).Clauses(clause.OnConflict{
		DoUpdates: clause.Assignments(map[string]interface{}{
			"name":        p.Name,
			"category_id": p.CategoryId,
			"brand_id":    p.BrandId,
			"description": p.Description,
			"is_new":      p.IsNew,
			"is_hot":      p.IsHot,
			"on_sale":     p.OnSale,
			"click":       p.Click,
			"sale":        p.Sale,
			"favorite":    p.Favorite,
			"mark_price":  p.MarkPrice,
			"shop_price":  p.ShopPrice,
			"picture":     p.Picture,
			"images":      p.Images,
			"update_at":   now,
		}),
	}).Create(&p)

	if result.Error != nil {
		return Products{}, result.Error
	}

	return p, nil
}

func (this *product) DeleteProduct(ctx context.Context, id int32) error {
	now := time.Now().UnixMilli()
	if err := this.db.WithContext(ctx).Where("id = ?", id).Updates(Products{DeleteAt: now}).Error; err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			return ErrRecordNotFound
		}
		return err
	}
	return nil
}

func (this *product) QueryProductInfoById(ctx context.Context, id int32) (Products, error) {
	var p Products
	result := this.db.WithContext(ctx).Where("id = ? AND delete_at = 0", id).First(&p)
	if result.Error != nil {
		return Products{}, result.Error
	}

	if result.RowsAffected == 0 {
		return Products{}, ErrRecordNotFound
	}

	return p, nil
}

func NewProductDao(db *gorm.DB, logger *zap.Logger) Product {
	return &product{
		db:     db,
		logger: logger,
	}
}
