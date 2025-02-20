package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"

	"github.com/Numsina/tkshop/app/products/domain"
	"github.com/Numsina/tkshop/app/products/repository/dao"
)

type Product interface {
	CraeteOrUpdateProduct(ctx context.Context, p domain.Products) (domain.Products, error)
	DeleteProduct(ctx context.Context, id, uid int32) error
	GetProductInfoById(ctx context.Context, id int32) (domain.Products, error)
	GetProductList(ctx context.Context, p domain.Products, pNum, pSize int32) ([]domain.Products, error)
	AddClick(ctx context.Context, id, uid int32) error
	AddFavorite(ctx context.Context, id, uid int32) error

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
	d dao.Product
}

func (r *product) GetBrandByUid(ctx context.Context, uid int32) ([]domain.Brands, error) {
	data, err := r.d.QueryBrandByUid(ctx, uid)
	if err != nil {
		return nil, err
	}
	var result []domain.Brands
	for _, v := range data {
		result = append(result, r.toBrandDomain(v))
	}
	return result, nil
}

func NewProductRepo(d dao.Product) Product {
	return &product{d: d}
}

/**
缓存待优化
*/

func (r *product) GetBrandList(ctx *gin.Context, num int32, size int32) ([]domain.Brands, error) {
	if num <= 0 {
		num = 1
	}
	if size < 20 {
		size = 20
	}
	data, err := r.d.QueryBrandList(ctx, num, size)
	if err != nil {
		return nil, err
	}
	var result []domain.Brands
	for _, v := range data {
		result = append(result, r.toBrandDomain(v))
	}
	return result, nil
}

func (r *product) DeleteBrand(ctx context.Context, id, uid int32) error {
	return r.d.DeleteBrand(ctx, id, uid)
}

func (r *product) UpdateBrand(ctx context.Context, c domain.Brands) error {
	return r.d.UpdateBrand(ctx, r.toBrandDao(c))
}

func (r *product) CreateBrand(ctx context.Context, c domain.Brands) (domain.Brands, error) {
	data, err := r.d.InsertBrand(ctx, r.toBrandDao(c))
	if err != nil {
		return domain.Brands{}, err
	}
	return r.toBrandDomain(data), nil
}

func (r *product) GetBrandByName(ctx context.Context, name string) (domain.Brands, error) {
	data, err := r.d.QueryBrandByName(ctx, name)
	if err != nil {
		return domain.Brands{}, err
	}
	return r.toBrandDomain(data), nil
}

func (r *product) GetBrandById(ctx context.Context, id int32) (domain.Brands, error) {
	data, err := r.d.QueryBrandById(ctx, id)
	if err != nil {
		return domain.Brands{}, err
	}
	return r.toBrandDomain(data), nil
}

func (r *product) GetCategoryList(ctx *gin.Context, num int32, size int32) ([]domain.Categorys, error) {
	if num <= 0 {
		num = 1
	}
	if size < 20 {
		size = 20
	}
	data, err := r.d.QueryCategoryList(ctx, num, size)
	if err != nil {
		return nil, err
	}
	var result []domain.Categorys
	for _, v := range data {
		result = append(result, r.toCategoryDomain(v))
	}
	return result, nil
}

func (r *product) DeleteCategory(ctx context.Context, id, uid int32) error {
	return r.d.DeleteCategory(ctx, id, uid)
}

func (r *product) UpdateCategory(ctx context.Context, c domain.Categorys) error {
	return r.d.UpdateCategory(ctx, r.toCategoryDao(c))
}
func (r *product) CreateCategory(ctx context.Context, c domain.Categorys) (domain.Categorys, error) {

	data, err := r.d.InsertCategory(ctx, r.toCategoryDao(c))
	if err != nil {
		return domain.Categorys{}, err
	}
	return r.toCategoryDomain(data), nil
}

func (r *product) GetCategoryByName(ctx context.Context, name string) (domain.Categorys, error) {
	data, err := r.d.QueryCategoryByName(ctx, name)
	if err != nil {
		return domain.Categorys{}, err
	}
	return r.toCategoryDomain(data), nil
}

func (r *product) GetCategoryById(ctx context.Context, id int32) (domain.Categorys, error) {
	data, err := r.d.QueryCategoryById(ctx, id)
	if err != nil {
		return domain.Categorys{}, err
	}
	return r.toCategoryDomain(data), nil
}

func (r *product) GetProductList(ctx context.Context, p domain.Products, pNum, pSize int32) ([]domain.Products, error) {
	if pNum <= 0 {
		pNum = 1
	}
	if pSize < 10 || pSize > 500 {
		pSize = 20
	}

	if p.CategoryName != "" {
		data, err := r.d.QueryCategoryByName(ctx, p.CategoryName)
		if err != nil {
			return nil, err
		}
		p.CategoryId = data.Id
	}

	if p.BrandName != "" {
		data, err := r.d.QueryBrandByName(ctx, p.BrandName)
		if err != nil {
			return nil, err
		}
		p.BrandId = data.Id
	}

	products, err := r.d.QueryProductList(ctx, r.toDao(p), pNum, pSize)
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, nil
	}
	var data []domain.Products
	for _, v := range products {
		data = append(data, r.toDomain(v))
	}
	return data, nil
}

func (r *product) AddClick(ctx context.Context, id, uid int32) error {
	return r.d.AddClick(ctx, id, uid)
}

func (r *product) AddFavorite(ctx context.Context, id, uid int32) error {
	return r.d.AddFavorite(ctx, id, uid)
}

func (r *product) CraeteOrUpdateProduct(ctx context.Context, p domain.Products) (domain.Products, error) {
	sn := fmt.Sprintf("%s:%d:%d:%d", p.Name, p.CategoryId, p.BrandId, p.Uid)
	if err := r.d.GetProductBySn(ctx, sn); err != nil {
		return domain.Products{}, err
	}
	p.Sn = sn
	data, err := r.d.UpsertProduct(ctx, r.toDao(p))

	return r.toDomain(data), err
}

func (r *product) DeleteProduct(ctx context.Context, id, uid int32) error {
	return r.d.DeleteProduct(ctx, id, uid)
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
		Uid:         p.Uid,
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
		Sn:          p.Sn,
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

func (r *product) toCategoryDomain(c dao.Categorys) domain.Categorys {
	data := domain.Categorys{
		Id:    c.Id,
		Name:  c.Name,
		Level: c.Level,
	}
	if c.ParentId.Valid {
		data.ParentId = c.ParentId.Int32
	}

	if c.RootId.Valid {
		data.RootId = c.RootId.Int32
	}

	return data
}

func (r *product) toCategoryDao(c domain.Categorys) dao.Categorys {
	data := dao.Categorys{
		Id:    c.Id,
		Name:  c.Name,
		Level: c.Level,
		Uid:   c.Uid,
		ParentId: sql.NullInt32{
			Int32: c.ParentId,
			Valid: c.ParentId == 0,
		},
		RootId: sql.NullInt32{
			Int32: c.RootId,
			Valid: c.RootId == 0,
		},
	}
	return data
}

func (r *product) toBrandDomain(b dao.Brands) domain.Brands {
	data := domain.Brands{
		Id:   b.Id,
		Name: b.Name,
		Logo: b.Logo,
	}
	return data
}

func (r *product) toBrandDao(b domain.Brands) dao.Brands {
	data := dao.Brands{
		Id:   b.Id,
		Name: b.Name,
		Logo: b.Logo,
		Uid:  b.Uid,
	}
	return data
}
