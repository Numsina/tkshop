package router

import (
	"github.com/gin-gonic/gin"

	"github.com/Numsina/tkshop/app/products/ioc"
)

func InitProductRouter(router *gin.RouterGroup) {
	route := ioc.InitProduct()
	productRouter := router.Group("/products")
	{
		productRouter.GET("/:id", route.GetProductsDetail)
		productRouter.POST("", route.Create)
		productRouter.PUT("", route.Update)
		productRouter.DELETE("", route.DeleteProduct)
		productRouter.GET("/list", route.GetProductsList)
		productRouter.POST("/like", route.AddFavorite)
	}

	categoryRouter := router.Group("/categorys")
	{
		categoryRouter.GET("/:id", route.GetCategoryById)
		categoryRouter.POST("", route.CreateCategory)
		categoryRouter.PUT("", route.UpdateCategory)
		categoryRouter.DELETE("", route.DeleteCategory)
		categoryRouter.GET("", route.GetCategoryByName)
		categoryRouter.POST("/list", route.GetCategoryList)
	}

	brandRouter := router.Group("/brands")
	{
		brandRouter.GET("/:id", route.GetBrandById)
		brandRouter.POST("", route.CreateBrand)
		brandRouter.PUT("", route.UpdateBrand)
		brandRouter.DELETE("", route.DeleteBrand)
		brandRouter.POST("/list", route.GetBrandList)
	}
}
