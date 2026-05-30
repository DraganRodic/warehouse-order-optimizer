package router

import (
	"github.com/DraganRodic/warehouse-order-optimizer/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/inventory/upload", handler.UploadInventory)
		api.POST("/inventory/import", handler.ImportProducts)

		api.POST("/orders/process", handler.ProcessOrder)

		api.GET("/products", handler.GetProducts)
		api.GET("/products/:id", handler.GetProductByID)
		api.POST("/products", handler.CreateProduct)
		api.PUT("/products/:id", handler.UpdateProduct)
		api.DELETE("/products/:id", handler.DeleteProduct)
	}

	return r
}
