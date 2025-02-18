package main

import (
	"github.com/Numsina/tkshop/app/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.InitRouter(r)

	r.Run(":9988")
}
