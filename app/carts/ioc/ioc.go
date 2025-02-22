package ioc

import (
	"github.com/Numsina/tkshop/app/carts/handler"
	"github.com/Numsina/tkshop/app/carts/initialize"
	"github.com/Numsina/tkshop/app/carts/repository"
	"github.com/Numsina/tkshop/app/carts/repository/dao"
)

func InitCart() *handler.CartHandler {
	initialize.InitConfig()
	logger := initialize.InitLogger()
	db := initialize.InitDB()
	dao.InitCartsTable(db)
	entity := dao.New(db, logger)
	cartRepo := repository.NewCartRepository(entity, logger)
	return handler.NewCartHandler(cartRepo)
}
