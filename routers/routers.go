package routers

import (
	"github.com/labstack/echo/v4"
	"qcafe/controllers"
)

func SetRouters(e *echo.Echo) {
	routers := e.Group("/api")

	routers.GET("/products", controllers.GetProducts)
	routers.GET("/products/:code", controllers.GetProductByCode)
	routers.POST("/products", controllers.CreateProduct)
	routers.PUT("/products/:code", controllers.UpdateProductByCode)
	routers.DELETE("/products/:code", controllers.DeleteProductByCode)

	routers.GET("/categories", controllers.GetCategories)
	routers.GET("/categories/:code", controllers.GetCategoryByCode)
	routers.POST("/categories", controllers.CreateCategory)
	routers.PUT("/categories/:code", controllers.UpdateCategoryByCode)
	routers.DELETE("/categories/:code", controllers.DeleteCategoryByCode)
}
