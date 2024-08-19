package router

import (
	"crud-go/config"
	"crud-go/handler"
	"crud-go/repository"
	"crud-go/service"

	"github.com/gin-gonic/gin"
)

func PostRouter(api *gin.RouterGroup) {
	productRepository := repository.NewProductRepository(config.DB)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	r := api.Group("/products")

	r.GET("/", productHandler.GetProducts)
	r.POST("/", productHandler.CreateProduct)
}