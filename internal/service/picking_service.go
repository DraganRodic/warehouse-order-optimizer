package service

import (
	"sort"

	"github.com/DraganRodic/warehouse-order-optimizer/internal/model"
)

func SortByLocation(products []model.Product) {

	sort.Slice(products, func(i, j int) bool {
		return products[i].Location < products[j].Location
	})
}