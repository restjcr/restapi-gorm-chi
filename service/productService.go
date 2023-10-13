package service

import (
	"github.com/restjcr/restapi-gorm-chi/model"
	"github.com/restjcr/restapi-gorm-chi/repository"
	"gorm.io/gorm"
	"log"
)

type ProductService struct {
	repository *repository.ProductRepository
}

func NewProductService(repository *repository.ProductRepository) *ProductService {
	return &ProductService{
		repository: repository,
	}
}

func (p *ProductService) GetAllProducts() (*[]model.Product, error) {
	var products []model.Product
	p.repository.Database.Find(&products)

	return &products, nil

}

func (p *ProductService) GetProduct(productId int) (*model.Product, error) {
	var product model.Product
	if err := p.repository.Database.First(&product, productId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("Item not found")
		}
		return nil, err
	}

	return &product, nil
}

func (p *ProductService) CreateProduct(product model.Product) error {
	err := p.repository.Database.Create(&product).Error

	if err != nil {
		log.Println("Cannot insert that record")
		return err
	}

	return nil

}

func (p *ProductService) UpdateProduct(updatedProduct model.Product, productId int) (*model.Product, error) {
	var actualProduct model.Product

	if err := p.repository.Database.First(&actualProduct, productId).Error; err != nil {
		log.Println("Item not found")
		return nil, err
	}

	actualProduct.Name = updatedProduct.Name
	actualProduct.Price = updatedProduct.Price

	if err := p.repository.Database.Save(&actualProduct).Error; err != nil {
		log.Println("Cannot update the item")
		return nil, err
	}

	return &actualProduct, nil

}

func (p *ProductService) DeleteProduct(productId int) error {
	var product model.Product

	if err := p.repository.Database.First(&product, productId).Error; err != nil {
		log.Println("Item not found")
		return err
	}

	if err := p.repository.Database.Delete(&product).Error; err != nil {
		log.Println("Cannot delete that product")
		return err
	}

	return nil
}
