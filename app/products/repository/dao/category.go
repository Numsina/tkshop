package dao

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func (this *product) QueryCategoryList(ctx context.Context, num, size int32) ([]Categorys, error) {
	var data []Categorys
	of := (num - 1) * size
	result := this.db.WithContext(ctx).Offset(int(of)).Limit(int(size)).Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("没有分类")
	}
	return data, nil
}

func (this *product) QueryCategoryById(ctx context.Context, id int32) (Categorys, error) {
	var data Categorys
	result := this.db.WithContext(ctx).First(&data, id)
	fmt.Println(result.Error)
	if result.Error != nil {
		return Categorys{}, result.Error
	}
	if result.RowsAffected == 0 {
		return Categorys{}, errors.New("查询的商品不存在")
	}
	return data, nil
}

func (this *product) QueryCategoryByName(ctx context.Context, name string) (Categorys, error) {
	var data Categorys
	err := this.db.WithContext(ctx).Where("name = ?", name).First(&data).Error
	return data, err
}

func (this *product) InsertCategory(ctx context.Context, c Categorys) (Categorys, error) {
	var cg Categorys
	if c.ParentId.Int32 != 0 {
		// 可以用定时任务，将商品分类id存入redis, 查询商品是否存在(当然添加商品不是一个并发高的行为， 缓不缓存无所谓)
		err := this.db.WithContext(ctx).Where("parentId = ?", c.ParentId).First(&cg).Error
		if err != nil {
			this.logger.Sugar().Infof("商品分类不存在, %d", c.ParentId)
			return Categorys{}, err
		}
	}

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
	err := this.db.WithContext(ctx).Where("id = ? AND uid = ?", c.Id, c.Uid).Updates(Categorys{
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

func (this *product) DeleteCategory(ctx context.Context, id, uid int32) error {

	result := this.db.WithContext(ctx).Model(&Categorys{}).Delete(Categorys{Id: id, Uid: uid})
	if result.Error != nil {
		this.logger.Sugar().Errorf("删除商品分类失败, 失败原因为：%s", result.Error)
		return result.Error
	}

	if result.RowsAffected != 1 {
		return ErrRecordNotFound
	}
	return nil
}
