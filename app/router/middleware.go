package router

import (
	"github.com/Numsina/tkshop/app/middlewares/trace"
	"github.com/Numsina/tkshop/pkg/metricsx/prome"
	"github.com/gin-gonic/gin"

	"github.com/Numsina/tkshop/app/middlewares"
	"github.com/Numsina/tkshop/app/user/initialize"
	"github.com/Numsina/tkshop/internal/metrcs"
)

func initMiddleWare(r *gin.Engine) {
	r.Use(middlewares.Cors())
	initialize.InitConfig()

	metrcs.InitPrometheus()
	jhl := middlewares.NewJWT([]byte(initialize.Conf.JwtInfo.Key))
	auth := middlewares.NewLoginJWTMiddleWareBuilder(jhl)

	metric := prome.NewMetrics("tkshop", "tkshkop-v01", "tkshop", "tkshop", "统计http请求响应时间")

	r.Use(
		//trace.Trace(),
		trace.NewZipKinMiddleware().Trace(),
		metric.Build(),
		auth.IngorePaths("/v1/users/signup").
			IngorePaths("/test/metric").
			IngorePaths("/v1/products/list").
			IngorePaths("/v1/users/login").IngorePaths("/swagger/index.html").Build(),
	)
}
