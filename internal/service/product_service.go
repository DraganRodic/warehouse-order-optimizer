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

func (s *ProductService) FindProductsBySKU(skus []string) ([]model.Product, error) {
	return s.repo.FindBySKU(skus)
}

func (s *ProductService) GetAllProducts() ([]model.Product, error) {
	return s.repo.FindAll()
}

func (s *ProductService) GetProductByID(id uint) (*model.Product, error) {
	return s.repo.FindByID(id)
}

func (s *ProductService) CreateProduct(product *model.Product) error {
	return s.repo.Create(product)
}

func (s *ProductService) UpdateProduct(product *model.Product) error {
	return s.repo.Update(product)
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.repo.Delete(id)
}