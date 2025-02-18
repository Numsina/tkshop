package router

import (
	"github.com/gin-gonic/gin"

	"github.com/Numsina/tkshop/app/middlewares"
	"github.com/Numsina/tkshop/app/user/ioc"
)

func InitUserRouter(router *gin.RouterGroup) {
	route := ioc.InitUser()
	userRouter := router.Group("/users").Use(middlewares.Cors())
	{
		userRouter.GET("/info", route.GetUserByEmail)
		userRouter.POST("/singup", route.SingUp)
		userRouter.POST("/update", route.Update)
		userRouter.POST("/login", route.Login)
		userRouter.DELETE("/delete", route.Delete)
	}
}
