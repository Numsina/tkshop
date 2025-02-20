package service

import (
	"context"

	"github.com/Numsina/tkshop/app/products/domain"
)

func (r *product) GetBrandById(ctx context.Context, id int32) (domain.Brands, error) {
	return r.repo.GetBrandById(ctx, id)
}

func (r *product) GetBrandByName(ctx context.Context, name string) (domain.Brands, error) {
	return r.repo.GetBrandByName(ctx, name)
}

func (r *product) CreateBrand(ctx context.Context, c domain.Brands) (domain.Brands, error) {
	return r.repo.CreateBrand(ctx, c)
}
func (r *product) UpdateBrand(ctx context.Context, c domain.Brands) error {
	return r.repo.UpdateBrand(ctx, c)
}
func (r *product) DeleteBrand(ctx context.Context, id int32) error {
	return r.repo.DeleteBrand(ctx, id)
}
