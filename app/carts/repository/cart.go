package repository

import (
	"context"

	"go.uber.org/zap"

	"github.com/Numsina/tkshop/app/carts/domain"
	"github.com/Numsina/tkshop/app/carts/repository/dao"
)

type Cart interface {
	CreateCarts(ctx context.Context, cart domain.Carts, uid int32) error
	DeleteCarts(ctx context.Context, pid, uid int32) error
	UpdateCarts(ctx context.Context, cart domain.Carts, uid int32) error
	ClearCarts(ctx context.Context, uid int32) error
	GetCartsInfo(ctx context.Context, uid int32) ([]domain.Carts, error)
}

type CartRepository struct {
	dao dao.Cart
	log *zap.Logger
}

func (c *CartRepository) CreateCarts(ctx context.Context, cart domain.Carts, uid int32) error {
	d := c.toDao(cart)
	d.UserID = uid
	return c.dao.InsertCarts(ctx, d)
}

func (c *CartRepository) DeleteCarts(ctx context.Context, pid, uid int32) error {
	return c.dao.DeleteCarts(ctx, pid, uid)
}

func (c *CartRepository) UpdateCarts(ctx context.Context, cart domain.Carts, uid int32) error {
	d := c.toDao(cart)
	d.UserID = uid
	return c.dao.UpdateCarts(ctx, d)
}

func (c *CartRepository) ClearCarts(ctx context.Context, uid int32) error {
	return c.dao.ClearCarts(ctx, uid)
}

func (c *CartRepository) GetCartsInfo(ctx context.Context, uid int32) ([]domain.Carts, error) {
	data, err := c.dao.QueryCartsInfo(ctx, uid)
	if err != nil {
		return nil, err
	}

	var ds []domain.Carts
	for _, v := range data {
		ds = append(ds, c.toDomain(v))
	}
	return ds, nil
}

func NewCartRepository(dao dao.Cart, log *zap.Logger) Cart {
	return &CartRepository{
		dao: dao,
		log: log,
	}
}

func (c *CartRepository) toDao(data domain.Carts) dao.Carts {
	return dao.Carts{
		Id:      data.Id,
		GoodsID: data.GoodsID,
		Nums:    data.Nums,
		Checked: data.Checked,
	}
}

func (c *CartRepository) toDomain(data dao.Carts) domain.Carts {
	return domain.Carts{
		Id:      data.Id,
		GoodsID: data.GoodsID,
		Nums:    data.Nums,
		Checked: data.Checked,
	}
}
