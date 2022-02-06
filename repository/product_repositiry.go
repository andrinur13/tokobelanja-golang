package repository

import (
	"tokobelanja-golang/model/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProductByID(ID int) (entity.Product, error)
	Update(newProduct entity.Product) (entity.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db}
}

func (r *productRepository) GetProductByID(ID int) (entity.Product, error) {
	productResult := entity.Product{}

	err := r.db.Where("id = ?", ID).Find(&productResult).Error

	if err != nil {
		return entity.Product{}, err
	}

	return productResult, nil
}

func (r *productRepository) Update(newProduct entity.Product) (entity.Product, error) {
	err := r.db.Where("id = ?", newProduct.ID).Updates(newProduct).Error

	if err != nil {
		return entity.Product{}, err
	}

	return newProduct, nil
}
