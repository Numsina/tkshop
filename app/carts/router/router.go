package router

import (
	"github.com/gin-gonic/gin"

	"github.com/Numsina/tkshop/app/carts/ioc"
)

func InitCartRouter(router *gin.RouterGroup) {
	route := ioc.InitCart()
	cartRouter := router.Group("/carts")
	{
		cartRouter.GET("", route.Get)
		cartRouter.POST("", route.Create)
		cartRouter.PUT("", route.Update)
		cartRouter.DELETE("", route.Delete)
		cartRouter.DELETE("/clear", route.Clear)
	}

}
