package router

import (
	"github.com/gin-gonic/gin"

	productR "github.com/Numsina/tkshop/app/products/router"
	userR "github.com/Numsina/tkshop/app/user/router"
)

func InitRouter(r *gin.Engine) {
	v1 := r.Group("/v1")
	userR.InitUserRouter(v1)
	productR.InitProductRouter(v1)
}
