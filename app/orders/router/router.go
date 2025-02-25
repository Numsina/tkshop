package router

import (
	"github.com/gin-gonic/gin"

	"github.com/Numsina/tkshop/app/orders/ioc"
)

func InitCartRouter(router *gin.RouterGroup) {
	route := ioc.InitOrders()
	orderRouter := router.Group("/orders")
	{
		orderRouter.GET("", route.GetByUid)
		orderRouter.POST("", route.Create)
		orderRouter.PUT("", route.Update)
		orderRouter.DELETE("", route.Delete)
		orderRouter.POST("/search", route.Search)
	}

}
