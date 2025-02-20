package service

import (
	"context"

	"github.com/Numsina/tkshop/app/products/domain"
)

func (r *product) GetCategoryById(ctx context.Context, id int32) (domain.Categorys, error) {
	return r.repo.GetCategoryById(ctx, id)
}

func (r *product) GetCategoryByName(ctx context.Context, name string) (domain.Categorys, error) {
	return r.repo.GetCategoryByName(ctx, name)
}

func (r *product) CreateCategory(ctx context.Context, c domain.Categorys) (domain.Categorys, error) {
	return r.repo.CreateCategory(ctx, c)
}
func (r *product) UpdateCategory(ctx context.Context, c domain.Categorys) error {
	return r.repo.UpdateCategory(ctx, c)
}
func (r *product) DeleteCategory(ctx context.Context, id int32) error {
	return r.repo.DeleteCategory(ctx, id)
}

