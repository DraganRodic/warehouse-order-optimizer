package repository

import (
	"github.com/DraganRodic/warehouse-order-optimizer/internal/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) CreateMany(products []model.Product) error {
	return r.db.Create(&products).Error
}