package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	cartR "github.com/Numsina/tkshop/app/carts/router"
	productR "github.com/Numsina/tkshop/app/products/router"
	userR "github.com/Numsina/tkshop/app/user/router"
)

func InitRouter(r *gin.Engine) {
	initMiddleWare(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("/v1")
	userR.InitUserRouter(v1)
	productR.InitProductRouter(v1)
	cartR.InitCartRouter(v1)
}
