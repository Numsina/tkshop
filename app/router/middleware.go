package router

import (
	"github.com/gin-gonic/gin"

	"github.com/Numsina/tkshop/app/middlewares"
	"github.com/Numsina/tkshop/app/user/initialize"
)

func initMiddleWare(r *gin.Engine) {
	r.Use(middlewares.Cors())
	initialize.InitConfig()
	jhl := middlewares.NewJWT([]byte(initialize.Conf.JwtInfo.Key))
	auth := middlewares.NewLoginJWTMiddleWareBuilder(jhl)
	r.Use(auth.IngorePaths("/v1/users/signup").
		IngorePaths("/v1/users/login").Build(),
	)
}
