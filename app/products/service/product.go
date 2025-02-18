package service

import (
	"context"

	"github.com/Numsina/tkshop/app/products/domain"
	"github.com/Numsina/tkshop/app/products/repository"
)

type Product interface {
	CraeteProduct(ctx context.Context, p domain.Products) (domain.Products, error)
	UpdateProduct(ctx context.Context, p domain.Products) (domain.Products, error)
	DeleteProduct(ctx context.Context, id int32) error
	GetProductInfoById(ctx context.Context, id int32) (domain.Products, error)
}

type product struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) Product {
	return &product{
		repo: repo,
	}
}

func (r *product) CraeteProduct(ctx context.Context, p domain.Products) (domain.Products, error) {
	return r.repo.CraeteOrUpdateProduct(ctx, p)
}

func (r *product) UpdateProduct(ctx context.Context, p domain.Products) (domain.Products, error) {
	return r.repo.CraeteOrUpdateProduct(ctx, p)
}

func (r *product) DeleteProduct(ctx context.Context, id int32) error {
	return r.repo.DeleteProduct(ctx, id)
}

func (r *product) GetProductInfoById(ctx context.Context, id int32) (domain.Products, error) {
	return r.repo.GetProductInfoById(ctx, id)
}
