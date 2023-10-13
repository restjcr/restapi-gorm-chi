package repository

import "gorm.io/gorm"

type ProductRepository struct {
	Database *gorm.DB
}

func NewProductRepository(database *gorm.DB) *ProductRepository {
	return &ProductRepository{
		Database: database,
	}
}
