package router

import (
	"github.com/gin-gonic/gin"

	"github.com/Numsina/tkshop/app/user/ioc"
)

func InitUserRouter(router *gin.RouterGroup) {
	route := ioc.InitUser()
	userRouter := router.Group("/users")
	{
		userRouter.GET("/info", route.GetUserByEmail)
		userRouter.POST("/signup", route.SignUp)
		userRouter.POST("/update", route.Update)
		userRouter.POST("/login", route.Login)
		userRouter.DELETE("/delete", route.Delete)
		userRouter.DELETE("/logout", route.Logout)
	}
}
