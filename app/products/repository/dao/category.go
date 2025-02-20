package dao

import (
	"context"
	"time"
)

func (this *product) QueryCategoryById(ctx context.Context, id int32) (Categorys, error) {
	var data Categorys
	err := this.db.WithContext(ctx).First(&data).Error
	return data, err
}

func (this *product) QueryCategoryByName(ctx context.Context, name string) (Categorys, error) {
	var data Categorys
	err := this.db.WithContext(ctx).Where("name = ?", name).First(&data).Error
	return data, err
}

func (this *product) InsertCategory(ctx context.Context, c Categorys) (Categorys, error) {
	now := time.Now().UnixMilli()
	c.CreateAt = now
	c.UpdateAt = now
	err := this.db.WithContext(ctx).Create(&c).Error
	if err != nil {
		this.logger.Sugar().Errorf("数据库发生插入错误, 错误原因：%s", err.Error())
		return Categorys{}, err
	}
	return c, err
}

func (this *product) UpdateCategory(ctx context.Context, c Categorys) error {
	err := this.db.WithContext(ctx).Where("id = ?", c.Id).Updates(Categorys{
		Name:     c.Name,
		Level:    c.Level,
		ParentId: c.ParentId,
		UpdateAt: time.Now().UnixMilli(),
	}).Error

	if err != nil {
		this.logger.Sugar().Infof("更商品分类数据失败, 失败原因：%s", err.Error())
		return err
	}
	return nil
}

func (this *product) DeleteCategory(ctx context.Context, id int32) error {
	result := this.db.WithContext(ctx).Model(&Categorys{}).Delete(Categorys{Id: id})
	if result.Error != nil {
		this.logger.Sugar().Errorf("删除商品分类失败, 失败原因为：%s", result.Error)
		return result.Error
	}

	if result.RowsAffected != 1 {
		return ErrRecordNotFound
	}
	return nil
}
