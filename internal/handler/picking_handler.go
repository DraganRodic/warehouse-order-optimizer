package handler

import (
	"net/http"
	"path/filepath"

	"github.com/DraganRodic/warehouse-order-optimizer/internal/database"
	"github.com/DraganRodic/warehouse-order-optimizer/internal/repository"
	"github.com/DraganRodic/warehouse-order-optimizer/internal/service"
	"github.com/gin-gonic/gin"
)

func ProcessOrder(c *gin.Context) {

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

	skus, err := service.ReadOrderFile(savePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	repo := repository.NewProductRepository(database.DB)
	productService := service.NewProductService(repo)

	products, err := productService.FindProductsBySKU(skus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	service.SortByLocation(products)

	c.JSON(http.StatusOK, gin.H{
		"count":    len(products),
		"products": products,
	})
}
