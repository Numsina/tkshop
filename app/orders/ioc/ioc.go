package ioc

import (
	"github.com/Numsina/tkshop/app/orders/handler"
	"github.com/Numsina/tkshop/app/orders/initialize"
	"github.com/Numsina/tkshop/app/orders/repository"
	"github.com/Numsina/tkshop/app/orders/repository/dao"
)

func InitOrders() *handler.OrderHandler {
	initialize.InitConfig()
	logger := initialize.InitLogger()
	db := initialize.InitDB()
	dao.InitOrderTable(db)
	entity := dao.NewOrder(db, logger)
	orderRepo := repository.NewOrder(entity, logger)
	return handler.NewOrderHandler(orderRepo)
}
