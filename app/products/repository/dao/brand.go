package dao

import (
	"context"
	"errors"
	"time"
)

func (this *product) QueryBrandList(ctx context.Context, num, size int32) ([]Brands, error) {
	var data []Brands
	of := (num - 1) * size
	result := this.db.WithContext(ctx).Offset(int(of)).Limit(int(size)).Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("没有品牌")
	}
	return data, nil
}

// QueryBrandByUid 查看旗下品牌
func (this *product) QueryBrandByUid(ctx context.Context, uid int32) ([]Brands, error) {
	var data []Brands
	result := this.db.WithContext(ctx).Where("uid = ?", uid).Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return data, nil
}

func (this *product) QueryBrandById(ctx context.Context, id int32) (Brands, error) {
	var data Brands
	err := this.db.WithContext(ctx).First(&data, id).Error
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

	// 按理说这里插入完成后应该向第三张表中插入, 但是等会做！！！！
	return b, err
}

func (this *product) UpdateBrand(ctx context.Context, b Brands) error {
	err := this.db.WithContext(ctx).Where("id = ? AND uid = ?", b.Id, b.Uid).Updates(Brands{
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

func (this *product) DeleteBrand(ctx context.Context, id, uid int32) error {
	result := this.db.WithContext(ctx).Model(&Brands{}).Delete(Brands{Id: id, Uid: uid})
	if result.Error != nil {
		this.logger.Sugar().Errorf("删除商品分类失败, 失败原因为：%s", result.Error)
		return result.Error
	}

	if result.RowsAffected != 1 {
		return ErrRecordNotFound
	}
	return nil
}
