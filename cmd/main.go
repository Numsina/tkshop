package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Numsina/tkshop/app/router"
)

// @title Swagger Example API
// @version 1.0
// @description 这是一个简单的 API 文档示例
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /v1
func main() {
	r := gin.Default()

	router.InitRouter(r)

	r.Run(":9988")
}
