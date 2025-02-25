package dao

import (
	"context"
	"fmt"
	"gorm.io/gorm/clause"
	"math/rand"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"

	carts "github.com/Numsina/tkshop/app/carts/repository/dao"
	products "github.com/Numsina/tkshop/app/products/repository/dao"
)

type Order interface {
	InsertOrder(ctx context.Context, od Orders) error
	UpdateOrder(ctx context.Context, od Orders) error
	DeleteOrder(ctx context.Context, sn string, uid int32) error
	SearchOrder(ctx context.Context, od Orders, page int, size int) ([]Orders, error)
	QueryOrderByUid(ctx context.Context, uid, page, size int32) ([]Orders, error)
}

var _ Order = &order{}

type order struct {
	db          *gorm.DB
	logger      *zap.Logger
	cartsClient carts.Cart
	goodsClient products.Product
}

func NewOrder(db *gorm.DB, logger *zap.Logger) Order {
	return &order{
		db:          db,
		logger:      logger,
		cartsClient: carts.New(db, logger),
		goodsClient: products.NewProductDao(db, logger),
	}
}

func Generate(uid int32) string {
	seed := rand.Intn(89) + 10
	now := time.Now()
	return fmt.Sprintf("%d%d%d%d%d%d%d%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), seed, uid)
}

func (o *order) InsertOrder(ctx context.Context, od Orders) error {
	// 查询用户要下单的商品
	selected, err := o.cartsClient.FindSelected(ctx, od.UserId, true)
	if err != nil {
		return err
	}

	// 去查询商品
	var ids []int32
	var productNums = make(map[int32]int32, len(selected))
	for _, v := range selected {
		ids = append(ids, v.GoodsID)
		productNums[v.GoodsID] = v.Nums
	}

	p, err := o.goodsClient.BatchProductsByIds(ctx, ids)
	if err != nil {
		return err
	}

	// 计算价格, 并将每个商品记录下来
	var totalPrice float32
	var goodsInfo []OrderGoods
	for _, v := range p {
		price := v.ShopPrice * float32(productNums[v.Id])
		totalPrice += price
		goodsInfo = append(goodsInfo, OrderGoods{
			GoodsId: v.Id,
			Price:   v.ShopPrice,
			Nums:    productNums[v.Id],
		})
		// 应该生成一个订单记录
	}

	// 扣减库存

	// 生成订单
	// 生成订单号
	od.OrderSn = Generate(od.UserId)
	now := time.Now().UnixMilli()
	od.CreateAt = now
	od.UpdateAt = now
	tx := o.db.WithContext(ctx).Begin()
	if err = tx.Create(&od).Error; err != nil {
		tx.Rollback()
		o.logger.Error("create order failed", zap.Error(err))
		return err
	}

	if err = tx.Create(&od).Error; err != nil {
		tx.Rollback()
	}

	for _, v := range goodsInfo {
		v.OrderSn = od.OrderSn
	}
	if err = tx.CreateInBatches(&goodsInfo, 100).Error; err != nil {
		tx.Rollback()
		o.logger.Error("create order failed", zap.Error(err))
		return err
	}

	// 清除购物车中的选中
	err = o.cartsClient.BatchDeleteCarts(ctx, ids, od.UserId)
	if err != nil {
		tx.Rollback()
		o.logger.Error("调用购物车批量删除失败", zap.Error(err))
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		o.logger.Error("commit order failed", zap.Error(err))
		return err
	}
	return nil
}

func (o *order) UpdateOrder(ctx context.Context, od Orders) error {
	err := o.db.WithContext(ctx).Where("order_sn = ?", od.OrderSn).Clauses(
		clause.OnConflict{DoUpdates: clause.Assignments(map[string]interface{}{
			"pay_type": od.PayType,
			"status":   od.Status,
			"pay_time": od.PayTime,
		})}).Error
	if err != nil {
		o.logger.Error("update order failed", zap.Error(err))
		return err
	}
	return nil
}

func (o *order) DeleteOrder(ctx context.Context, sn string, uid int32) error {
	tx := o.db.WithContext(ctx).Begin()

	if err := tx.Where("order_sn = ? AND user_id = ?", sn, uid).Delete(&Orders{}).Error; err != nil {
		tx.Rollback()
		o.logger.Error("delete order failed", zap.Error(err))
		return err
	}

	if err := tx.Where("order_sn = ?", sn).Delete(&OrderGoods{}).Error; err != nil {
		tx.Rollback()
		o.logger.Error("delete order failed", zap.Error(err))
		return err
	}
	return nil
}

func (o *order) SearchOrder(ctx context.Context, od Orders, page int, size int) ([]Orders, error) {
	tx := o.db.WithContext(ctx)

	if od.OrderSn != "" {
		tx = tx.Where("order_sn = ?", od.OrderSn)
	}

	if od.Status != 0 && od.Status < 5 {
		tx = tx.Where("status = ?", od.Status)
	}

	if od.PayType != "" {
		tx = tx.Where("pay_type = ?", od.PayType)
	}

	if od.PayTime != 0 {
		tx = tx.Where("pay_time = ?", od.PayTime)
	}
	var ods []Orders
	of := (page - 1) * size
	if err := tx.Offset(of).Limit(page).Find(&ods).Error; err != nil {
		o.logger.Error("search order failed", zap.Error(err))
		return nil, err
	}
	return ods, nil
}

func (o *order) QueryOrderByUid(ctx context.Context, uid, page, size int32) ([]Orders, error) {
	var data []Orders
	of := (page - 1) * size
	if err := o.db.WithContext(ctx).Where("user_id = ?", uid).Offset(int(of)).Limit(int(size)).Find(&data).Error; err != nil {
		o.logger.Error("query order failed", zap.Error(err))
		return nil, err
	}
	return data, nil
}
