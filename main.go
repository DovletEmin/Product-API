// @title Product Management API
// @version 1.0
// @description Simple product management service using Gin.
// @host localhost:8080
// @BasePath /api

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/you/product-api/handlers"
	"github.com/you/product-api/store"

	
    docs "github.com/you/product-api/docs"
    "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
)

func main() {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Временное (in-memory) хранилище
	st := store.NewInMemoryStore()

	// Ручки (handlers)
	h := handlers.NewProductHandler(st)

	// Роуты
	api := r.Group("/api")
	{
		products := api.Group("/products")
		{
			products.GET("", h.ListProducts)        // GET /api/products
			products.POST("", h.CreateProduct)      // POST /api/products
			products.GET("/:id", h.GetProduct)      // GET /api/products/:id
			products.PUT("/:id", h.UpdateProduct)   // PUT /api/products/:id
			products.DELETE("/:id", h.DeleteProduct)// DELETE /api/products/:id
		}
	}

	// Запуск на :8080
	r.Run(":8080")
}
