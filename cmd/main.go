package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Numsina/tkshop/app/router"
)

func main() {
	r := gin.Default()

	router.InitRouter(r)

	r.Run(":9988")
}
