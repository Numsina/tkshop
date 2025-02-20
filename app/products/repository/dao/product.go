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
	QueryProductList(ctx context.Context, p Products, pNum, pSize int32) ([]Products, error)
	AddClick(ctx context.Context, id int32) error
	AddFavorite(ctx context.Context, id int32) error
	QueryCategoryById(ctx context.Context, id int32) (Categorys, error)
	QueryCategoryByName(ctx context.Context, name string) (Categorys, error)
	InsertCategory(ctx context.Context, c Categorys) (Categorys, error)
	UpdateCategory(ctx context.Context, c Categorys) error
	DeleteCategory(ctx context.Context, id int32) error
	QueryBrandById(ctx context.Context, id int32) (Brands, error)
	QueryBrandByName(ctx context.Context, name string) (Brands, error)
	InsertBrand(ctx context.Context, c Brands) (Brands, error)
	UpdateBrand(ctx context.Context, c Brands) error
	DeleteBrand(ctx context.Context, id int32) error
}

var _ Product = &product{}

type product struct {
	db     *gorm.DB
	logger *zap.Logger
}

func (this *product) QueryProductList(ctx context.Context, p Products, pNum, pSize int32) ([]Products, error) {
	state := this.db.WithContext(ctx).Model(&Products{})

	if p.BrandId > 0 {
		state = state.Where("brand_id = ?", p.BrandId)
	}

	if p.IsNew {
		state = state.Where("is_new = ?", p.IsNew)
	}

	if p.IsHot {
		state = state.Where("is_hot = ?", p.IsHot)
	}

	if p.MarkPrice > 0 {
		state = state.Where("mark_price >= ?", p.MarkPrice)
	}

	if p.ShopPrice > 0 {
		state = state.Where("shop_price = ?", p.ShopPrice)
	}

	if p.OnSale {
		state = state.Where("on_sale = ?", p.OnSale)
	}

	if p.Sale > 0 {
		state = state.Where("sale > ?", p.Sale)
	}
	type categoryIds struct {
		ID int32 `json:"id"`
	}
	var categorys []categoryIds
	if p.CategoryId > 0 {
		var category Categorys
		result := this.db.WithContext(ctx).First(&category)
		if result.Error != nil {
			this.logger.Sugar().Infof("数据库错误, 原因：%s", result.Error.Error())
			return nil, result.Error
		}

		if result.RowsAffected != 1 {
			return nil, errors.New("商品分类不存在")
		}

		if category.ParentId == 0 {
			result = this.db.WithContext(ctx).Model(&Categorys{}).Select("id").Where("root_id = ? AND level = 3", category.RootId, category.ParentId).Find(&categorys)
		} else {
			result = this.db.WithContext(ctx).Model(&Categorys{}).Select("id").Where("root_id = ? parentId = ? AND level = 3", category.RootId, category.ParentId).Find(&categorys)
		}

		if result.Error != nil {
			return nil, errors.New("该商品分类不存在商品")
		}

		if result.RowsAffected == 0 {
			return nil, errors.New("该商品分类不存在商品")
		}

		var ids []interface{}
		for _, v := range categorys {
			ids = append(ids, v.ID)
		}
		state = state.Where("category_id in ?", ids...)
	}

	off := (pNum-1)*pSize + pSize

	var ps []Products
	if err := state.Offset(int(off)).Limit(int(pSize)).Find(&ps).Error; err != nil {
		return nil, err
	}
	return ps, nil
}

func (this *product) UpsertProduct(ctx context.Context, p Products) (Products, error) {
	tx := this.db.WithContext(ctx).Begin()

	now := time.Now().UnixMilli()
	p.CreateAt = now
	p.UpdateAt = now
	result := tx.Model(&Products{}).Where("id = ?", p.Id).Clauses(clause.OnConflict{
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
		tx.Rollback()
		return Products{}, result.Error
	}

	// 查询商品的分类及品牌是否存在
	if err := tx.Model(&Categorys{}).Where("id = ?", p.CategoryId).Error; err != nil {
		tx.Rollback()
		return Products{}, err
	}

	if err := tx.Model(&Brands{}).Where("id = ?", p.BrandId).Error; err != nil {
		tx.Rollback()
		return Products{}, err
	}
	err := tx.Commit().Error
	if err != nil {
		for i := 0; i < 3; i++ {
			if err = tx.Commit().Error; err == nil {
				break
			}
		}
		if err != nil {
			tx.Rollback()
		}
	}

	return p, err
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

func (this *product) AddClick(ctx context.Context, id int32) error {
	if err := this.db.WithContext(ctx).Where("id = ?", id).Update("click", gorm.Expr("click + 1")).Error; err != nil {
		return err
	}
	return nil
}

func (this *product) AddFavorite(ctx context.Context, id int32) error {
	if err := this.db.WithContext(ctx).Where("id = ?", id).Update("favorite", gorm.Expr("favorite + 1")).Error; err != nil {
		return err
	}
	return nil
}

func NewProductDao(db *gorm.DB, logger *zap.Logger) Product {
	return &product{
		db:     db,
		logger: logger,
	}
}
