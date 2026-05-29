package service

import (
	"strings"

	"github.com/DraganRodic/warehouse-order-optimizer/internal/model"
	"github.com/xuri/excelize/v2"
)

func ReadInventoryFile(filePath string) ([]model.Product, error) {

	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = f.Close()
	}()

	sheetName := f.GetSheetName(0)

	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, err
	}

	var products []model.Product

	for i, row := range rows {

		if i < 3 {
			continue
		}

		if len(row) < 2 {
			continue
		}

		sku := strings.TrimSpace(row[0])
		location := strings.TrimSpace(row[1])

		if sku == "" || location == "" {
			continue
		}

		products = append(products, model.Product{
			SKU:      sku,
			Location: location,
		})
	}

	return products, nil
}