package service

import (
	"github.com/DraganRodic/warehouse-order-optimizer/internal/model"
	"github.com/DraganRodic/warehouse-order-optimizer/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) ImportProducts(products []model.Product) error {
	return s.repo.CreateMany(products)
}