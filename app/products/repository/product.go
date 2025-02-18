package repository

import (
	"context"

	"github.com/Numsina/tkshop/app/products/domain"
	"github.com/Numsina/tkshop/app/products/repository/dao"
)

type Product interface {
	CraeteOrUpdateProduct(ctx context.Context, p domain.Products) (domain.Products, error)
	DeleteProduct(ctx context.Context, id int32) error
	GetProductInfoById(ctx context.Context, id int32) (domain.Products, error)
}

type product struct {
	d dao.Product
}

func NewProductRepo(d dao.Product) Product {
	return &product{d: d}
}

func (r *product) CraeteOrUpdateProduct(ctx context.Context, p domain.Products) (domain.Products, error) {
	data, err := r.d.UpsertProduct(ctx, r.toDao(p))
	return r.toDomain(data), err
}

func (r *product) DeleteProduct(ctx context.Context, id int32) error {
	return r.d.DeleteProduct(ctx, id)
}

func (r *product) GetProductInfoById(ctx context.Context, id int32) (domain.Products, error) {
	result, err := r.d.QueryProductInfoById(ctx, id)
	return r.toDomain(result), err
}

func (r *product) toDao(p domain.Products) dao.Products {
	return dao.Products{
		Id:          p.Id,
		Name:        p.Name,
		CategoryId:  p.CategoryId,
		BrandId:     p.BrandId,
		Description: p.Description,
		IsNew:       p.IsNew,
		IsHot:       p.IsHot,
		OnSale:      p.OnSale,
		Click:       p.Click,
		Sale:        p.Sale,
		Favorite:    p.Favorite,
		MarkPrice:   p.MarkPrice,
		ShopPrice:   p.ShopPrice,
		Picture:     p.Picture,
		Images:      p.Images,
	}
}

func (r *product) toDomain(p dao.Products) domain.Products {
	return domain.Products{
		Id:          p.Id,
		Name:        p.Name,
		CategoryId:  p.CategoryId,
		BrandId:     p.BrandId,
		Description: p.Description,
		IsNew:       p.IsNew,
		IsHot:       p.IsHot,
		OnSale:      p.OnSale,
		Click:       p.Click,
		Sale:        p.Sale,
		Favorite:    p.Favorite,
		MarkPrice:   p.MarkPrice,
		ShopPrice:   p.ShopPrice,
		Picture:     p.Picture,
		Images:      p.Images,
	}
}
