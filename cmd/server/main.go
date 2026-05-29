package main

import (
	"github.com/DraganRodic/warehouse-order-optimizer/internal/database"
	"github.com/DraganRodic/warehouse-order-optimizer/internal/model"
	"github.com/DraganRodic/warehouse-order-optimizer/internal/router"
)

func main() {

	db := database.ConnectDB()

	db.AutoMigrate(&model.Product{})

	r := router.SetupRouter()

	r.Run(":8080")
}