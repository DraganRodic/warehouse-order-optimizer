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
	}

	return r
}
