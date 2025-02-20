package service

import (
	"context"
	"github.com/gin-gonic/gin"

	"github.com/Numsina/tkshop/app/products/domain"
	"github.com/Numsina/tkshop/app/products/repository"
)

type Product interface {
	CraeteProduct(ctx context.Context, p domain.Products) (domain.Products, error)
	UpdateProduct(ctx context.Context, p domain.Products) (domain.Products, error)
	DeleteProduct(ctx context.Context, id, uid int32) error
	GetProductInfoById(ctx context.Context, id int32) (domain.Products, error)
	GetProductList(ctx context.Context, p domain.Products, pNum, pSize int32) ([]domain.Products, error)
	IncreateClick(ctx context.Context, id, uid int32) error
	IncreateFavorite(ctx context.Context, id, uid int32) error

	GetCategoryById(ctx context.Context, id int32) (domain.Categorys, error)
	GetCategoryByName(ctx context.Context, name string) (domain.Categorys, error)
	CreateCategory(ctx context.Context, c domain.Categorys) (domain.Categorys, error)
	UpdateCategory(ctx context.Context, c domain.Categorys) error
	DeleteCategory(ctx context.Context, id, uid int32) error
	GetCategoryList(ctx *gin.Context, num int32, size int32) ([]domain.Categorys, error)
	GetBrandById(ctx context.Context, id int32) (domain.Brands, error)
	GetBrandByName(ctx context.Context, name string) (domain.Brands, error)
	CreateBrand(ctx context.Context, c domain.Brands) (domain.Brands, error)
	UpdateBrand(ctx context.Context, c domain.Brands) error
	DeleteBrand(ctx context.Context, id, uid int32) error
	GetBrandList(ctx *gin.Context, num int32, size int32) ([]domain.Brands, error)
	GetBrandByUid(ctx context.Context, uid int32) ([]domain.Brands, error)
}

var _ Product = &product{}

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

func (r *product) DeleteProduct(ctx context.Context, id, uid int32) error {
	return r.repo.DeleteProduct(ctx, id, uid)
}

func (r *product) GetProductInfoById(ctx context.Context, id int32) (domain.Products, error) {
	return r.repo.GetProductInfoById(ctx, id)
}

func (r *product) GetProductList(ctx context.Context, p domain.Products, pNum, pSize int32) ([]domain.Products, error) {
	return r.repo.GetProductList(ctx, p, pNum, pSize)
}

func (r *product) IncreateClick(ctx context.Context, id, uid int32) error {
	return r.repo.AddClick(ctx, id, uid)
}

func (r *product) IncreateFavorite(ctx context.Context, id, uid int32) error {
	return r.repo.AddFavorite(ctx, id, uid)
}
