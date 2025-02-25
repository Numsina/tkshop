package dao

import (
	"context"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Cart interface {
	InsertCarts(ctx context.Context, cart Carts) error
	DeleteCarts(ctx context.Context, pid, uid int32) error
	UpdateCarts(ctx context.Context, cart Carts) error
	ClearCarts(ctx context.Context, uid int32) error
	QueryCartsInfo(ctx context.Context, uid int32) ([]Carts, error)
	FindSelected(ctx context.Context, uid int32, checked bool) ([]Carts, error)
	BatchDeleteCarts(ctx context.Context, ids []int32, uid int32) error
}

var _ Cart = &cart{}

type cart struct {
	db     *gorm.DB
	logger *zap.Logger
}

func New(db *gorm.DB, logger *zap.Logger) Cart {
	return &cart{
		db:     db,
		logger: logger,
	}
}

func (c *cart) BatchDeleteCarts(ctx context.Context, ids []int32, uid int32) error {
	err := c.db.WithContext(ctx).Where("id in (?) AND uid = ?", ids, uid).Delete(&Carts{}).Error
	if err != nil {
		c.logger.Error("failed to delete cart by id", zap.Error(err))
		return err
	}
	return nil
}

func (c *cart) FindSelected(ctx context.Context, uid int32, checked bool) ([]Carts, error) {
	var data []Carts
	err := c.db.WithContext(ctx).Where("user_id = ? AND checked = ?", uid, checked).First(&data).Error
	if err != nil {
		c.logger.Error("find selected cart failed", zap.Error(err))
		return nil, err
	}
	return data, nil
}

func (c *cart) InsertCarts(ctx context.Context, cart Carts) error {
	// nums不能小于1
	// goodsID 一定要有对应的商品

	now := time.Now().UnixMilli()
	cart.CreateAt = now
	cart.UpdateAt = now
	result := c.db.WithContext(ctx).Create(&cart)
	if result.Error != nil {
		c.logger.Error("insert cart error", zap.Error(result.Error))
		return result.Error
	}
	return nil
}

func (c *cart) DeleteCarts(ctx context.Context, pid, uid int32) error {
	err := c.db.WithContext(ctx).Where(Carts{GoodsID: pid, UserID: uid}).Delete(&Carts{}).Error
	if err != nil {
		c.logger.Error("delete cart error", zap.Error(err))
		return err
	}
	return nil
}

func (c *cart) UpdateCarts(ctx context.Context, cart Carts) error {
	err := c.db.WithContext(ctx).Where("user_id = ? AND goods_id = ?", cart.UserID, cart.GoodsID).Clauses(clause.OnConflict{DoUpdates: clause.Assignments(map[string]interface{}{
		"nums":      cart.Nums,
		"checked":   cart.Checked,
		"update_at": time.Now().UnixMilli(),
	})}).Create(&cart).Error
	if err != nil {
		c.logger.Error("update cart error", zap.Error(err))
		return err
	}
	return nil
}

func (c *cart) ClearCarts(ctx context.Context, uid int32) error {
	err := c.db.WithContext(ctx).Where(Carts{UserID: uid}).Delete(&Carts{}).Error
	if err != nil {
		c.logger.Error("clear cart error", zap.Error(err))
		return err
	}
	return nil
}

func (c *cart) QueryCartsInfo(ctx context.Context, uid int32) ([]Carts, error) {
	// 做好分页
	var data []Carts
	err := c.db.WithContext(ctx).Where("user_id = ?", uid).Find(&data).Error
	if err != nil {
		c.logger.Error("query cart info error", zap.Error(err))
		return nil, err
	}
	return data, nil
}
