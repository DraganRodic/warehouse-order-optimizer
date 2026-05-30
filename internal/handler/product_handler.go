package handler

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/DraganRodic/warehouse-order-optimizer/internal/database"
	"github.com/DraganRodic/warehouse-order-optimizer/internal/model"
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

func GetProducts(c *gin.Context) {

	repo := repository.NewProductRepository(database.DB)
	productService := service.NewProductService(repo)

	products, err := productService.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count":    len(products),
		"products": products,
	})
}

func GetProductByID(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid product id",
		})
		return
	}

	repo := repository.NewProductRepository(database.DB)
	productService := service.NewProductService(repo)

	product, err := productService.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "product not found",
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {

	var product model.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	repo := repository.NewProductRepository(database.DB)
	productService := service.NewProductService(repo)

	if err := productService.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid product id",
		})
		return
	}

	var product model.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	product.ID = uint(id)

	repo := repository.NewProductRepository(database.DB)
	productService := service.NewProductService(repo)

	if err := productService.UpdateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid product id",
		})
		return
	}

	repo := repository.NewProductRepository(database.DB)
	productService := service.NewProductService(repo)

	if err := productService.DeleteProduct(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "product deleted successfully",
	})
}