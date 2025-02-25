package repository

import (
	"context"
	"go.uber.org/zap"

	"github.com/Numsina/tkshop/app/orders/domain"
	"github.com/Numsina/tkshop/app/orders/repository/dao"
)

type Order interface {
	SaveOrder(ctx context.Context, od domain.Orders, uid int32) error
	UpdateOrder(ctx context.Context, od domain.Orders, uid int32) error
	DeleteOrder(ctx context.Context, sn string, uid int32) error
	GetOrder(ctx context.Context, od domain.Orders, page int, size int) ([]domain.Orders, error)
	GetOrderByUid(ctx context.Context, uid, page, size int32) ([]domain.Orders, error)
}

type order struct {
	d      dao.Order
	logger *zap.Logger
}

func NewOrder(d dao.Order, logger *zap.Logger) Order {
	return &order{
		d:      d,
		logger: logger,
	}
}

func (o *order) SaveOrder(ctx context.Context, od domain.Orders, uid int32) error {
	data := o.toDao(od)
	data.UserId = uid
	return o.d.InsertOrder(ctx, data)
}

func (o *order) UpdateOrder(ctx context.Context, od domain.Orders, uid int32) error {
	data := o.toDao(od)
	data.UserId = uid
	return o.d.InsertOrder(ctx, data)
}

func (o *order) DeleteOrder(ctx context.Context, sn string, uid int32) error {
	return o.d.DeleteOrder(ctx, sn, uid)
}

func (o *order) GetOrder(ctx context.Context, od domain.Orders, page int, size int) ([]domain.Orders, error) {
	data := o.toDao(od)
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	ds, err := o.d.SearchOrder(ctx, data, page, size)
	if err != nil {
		return nil, err
	}
	var orders []domain.Orders
	for _, v := range ds {
		orders = append(orders, o.toDomain(v))
	}
	return orders, nil
}

func (o *order) GetOrderByUid(ctx context.Context, uid, page, size int32) ([]domain.Orders, error) {
	data, err := o.d.QueryOrderByUid(ctx, uid, page, size)
	if err != nil {
		return nil, err
	}
	var orders []domain.Orders
	for _, v := range data {
		orders = append(orders, o.toDomain(v))
	}
	return orders, nil
}

func (o *order) toDao(data domain.Orders) dao.Orders {
	return dao.Orders{
		OrderSn: data.OrderSn,
		PayType: data.PayType,
		Status:  data.Status,
		PayTime: data.PayTime,
		Address: data.Address,
		Phone:   data.Phone,
	}
}

func (o *order) toDomain(data dao.Orders) domain.Orders {
	return domain.Orders{
		OrderSn: data.OrderSn,
		PayType: data.PayType,
		Status:  data.Status,
		PayTime: data.PayTime,
		Address: data.Address,
		Phone:   data.Phone,
	}
}
