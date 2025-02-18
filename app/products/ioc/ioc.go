package ioc

import (
	"github.com/Numsina/tkshop/app/products/handler"
	"github.com/Numsina/tkshop/app/products/initialize"
	"github.com/Numsina/tkshop/app/products/repository"
	"github.com/Numsina/tkshop/app/products/repository/dao"
	"github.com/Numsina/tkshop/app/products/service"
)

func InitProduct() *handler.ProductHandler {
	initialize.InitConfig()
	logger := initialize.InitLogger()
	db := initialize.InitDB()
	dao.InitTable(db)
	entity := dao.NewProductDao(db, logger)
	productRepo := repository.NewProductRepo(entity)
	productSvc := service.NewProductService(productRepo)
	return handler.NewProductHandler(productSvc, logger)
}
