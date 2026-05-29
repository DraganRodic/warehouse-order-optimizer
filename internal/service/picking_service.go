package service

import (
	"sort"
	"strconv"
	"strings"

	"github.com/DraganRodic/warehouse-order-optimizer/internal/model"
)

func SortByLocation(products []model.Product) {

	sort.Slice(products, func(i, j int) bool {

		locI := products[i].Location
		locJ := products[j].Location

		numI, _ := strconv.Atoi(strings.TrimRight(locI, "LR"))
		numJ, _ := strconv.Atoi(strings.TrimRight(locJ, "LR"))

		if numI == numJ {
			return locI < locJ
		}

		return numI < numJ
	})
}