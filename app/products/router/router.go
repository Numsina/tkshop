package router

import (
	"github.com/gin-gonic/gin"

	"github.com/Numsina/tkshop/app/middlewares"
	"github.com/Numsina/tkshop/app/products/ioc"
)

func InitProductRouter(router *gin.RouterGroup) {
	route := ioc.InitProduct()
	productRouter := router.Group("/products").Use(middlewares.Cors())
	{
		productRouter.GET("/:id", route.GetProductsDetail)
		productRouter.POST("", route.Create)
		productRouter.PUT("", route.Update)
		productRouter.DELETE("", route.DeleteProduct)
	}
}
