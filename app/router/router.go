package router

import (
	cartR "github.com/Numsina/tkshop/app/carts/router"
	orderR "github.com/Numsina/tkshop/app/orders/router"
	productR "github.com/Numsina/tkshop/app/products/router"
	userR "github.com/Numsina/tkshop/app/user/router"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	initMiddleWare(r)
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("/v1")
	userR.InitUserRouter(v1)
	productR.InitProductRouter(v1)
	cartR.InitCartRouter(v1)
	orderR.InitCartRouter(v1)
}
