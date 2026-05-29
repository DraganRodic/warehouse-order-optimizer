package handler

import (
	"net/http"
	"path/filepath"

	"github.com/DraganRodic/warehouse-order-optimizer/internal/database"
	"github.com/DraganRodic/warehouse-order-optimizer/internal/repository"
	"github.com/DraganRodic/warehouse-order-optimizer/internal/service"
	"github.com/gin-gonic/gin"
)

func ImportProducts(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	savePath := filepath.Join("uploads", file.Filename)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	products, err := service.ReadInventoryFile(savePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	repo := repository.NewProductRepository(database.DB)
	productService := service.NewProductService(repo)

	if err := productService.ImportProducts(products); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"imported": len(products),
	})
}
