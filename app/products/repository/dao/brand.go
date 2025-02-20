package dao

import (
	"context"
	"time"
)

func (this *product) QueryBrandById(ctx context.Context, id int32) (Brands, error) {
	var data Brands
	err := this.db.WithContext(ctx).First(&data).Error
	return data, err
}

func (this *product) QueryBrandByName(ctx context.Context, name string) (Brands, error) {
	var data Brands
	err := this.db.WithContext(ctx).Where("name = ?", name).First(&data).Error
	return data, err
}

func (this *product) InsertBrand(ctx context.Context, b Brands) (Brands, error) {
	now := time.Now().UnixMilli()
	b.CreateAt = now
	b.UpdateAt = now
	err := this.db.WithContext(ctx).Create(&b).Error
	if err != nil {
		this.logger.Sugar().Errorf("数据库发生插入错误, 错误原因：%s", err.Error())
		return Brands{}, err
	}
	return b, err
}

func (this *product) UpdateBrand(ctx context.Context, b Brands) error {
	err := this.db.WithContext(ctx).Where("id = ?", b.Id).Updates(Brands{
		Name:     b.Name,
		Logo:     b.Logo,
		UpdateAt: time.Now().UnixMilli(),
	}).Error

	if err != nil {
		this.logger.Sugar().Infof("更商品分类数据失败, 失败原因：%s", err.Error())
		return err
	}
	return nil
}

func (this *product) DeleteBrand(ctx context.Context, id int32) error {
	result := this.db.WithContext(ctx).Model(&Brands{}).Delete(Brands{Id: id})
	if result.Error != nil {
		this.logger.Sugar().Errorf("删除商品分类失败, 失败原因为：%s", result.Error)
		return result.Error
	}

	if result.RowsAffected != 1 {
		return ErrRecordNotFound
	}
	return nil
}
