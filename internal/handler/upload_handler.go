package handler

import (
	"net/http"
	"path/filepath"

	"github.com/DraganRodic/warehouse-order-optimizer/internal/service"
	"github.com/gin-gonic/gin"
)

func UploadInventory(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	savePath := filepath.Join("uploads", file.Filename)

	err = c.SaveUploadedFile(file, savePath)
	if err != nil {
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

	limit := 10

	if len(products) < limit {
		limit = len(products)
	}

	c.JSON(http.StatusOK, gin.H{
		"count":    len(products),
		"products": products[:limit],
	})
}
