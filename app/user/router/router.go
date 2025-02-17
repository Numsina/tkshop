package router

import "github.com/gin-gonic/gin"

func InitUserRouter(router gin.RouterGroup) {
	userRouter := router.Group("users")
	{
		userRouter.GET("")
	}
}
