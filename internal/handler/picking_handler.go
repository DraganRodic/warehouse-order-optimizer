package handler

import (
	"net/http"
	"sort"

	"github.com/DraganRodic/warehouse-order-optimizer/internal/database"
	"github.com/DraganRodic/warehouse-order-optimizer/internal/model"
	"github.com/DraganRodic/warehouse-order-optimizer/internal/repository"
	"github.com/DraganRodic/warehouse-order-optimizer/internal/service"
	"github.com/gin-gonic/gin"
)

func ProcessOrder(c *gin.Context) {

	var req model.OrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	repo := repository.NewProductRepository(database.DB)
	productService := service.NewProductService(repo)

	products, err := productService.FindProductsBySKU(req.SKUs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	sort.Slice(products, func(i, j int) bool {
		return products[i].Location < products[j].Location
	})

	c.JSON(http.StatusOK, gin.H{
		"count":    len(products),
		"products": products,
	})
}