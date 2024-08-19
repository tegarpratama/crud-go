package repository

import (
	"crud-go/dto"
	"crud-go/entity"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface{
	TotalData(params *dto.FilterParams, categoryId string) (int64, error)
	FindAll(params *dto.FilterParams, categoryId string) (*[]dto.ProductResponse, error)
	CategoryExist(id uuid.UUID) error
	Create(product *entity.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository (db *gorm.DB) *productRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) TotalData(params *dto.FilterParams, categoryId string) (int64, error) {
	var count int64
	query := r.db.Model(&entity.Product{})

	if categoryId != "" {
		query = query.Where("category_id = ?", categoryId)
	}

	if params.Search != "" {
		searchPattern := fmt.Sprintf("%%%s%%", params.Search)
		query = query.Where("name LIKE ? OR description LIKE ?", searchPattern, searchPattern)
	}

	err := query.Count(&count)
	if err != nil {
		return count, err.Error
	}

	return count, nil
}

func (r *productRepository) FindAll(params *dto.FilterParams, categoryId string) (*[]dto.ProductResponse, error) {
	var product []dto.ProductResponse
	tx := r.db.Begin()

	query := tx.Model(&entity.Product{}).
		Select("id, name , description, category_id, DATE_FORMAT(created_at, '%Y-%m-%d %H:%i:%s') as created_at").
		Preload("Category")

	if categoryId != "" {
		query = query.Where("category_id = ?", categoryId)
	}

	if params.Search != "" {
		searchPattern := fmt.Sprintf("%%%s%%", params.Search)
		query = query.Where("name LIKE ? OR description LIKE ?", searchPattern, searchPattern)
	}

	err := query.Preload("Category").Order("created_at desc").Offset(params.Offset).Limit(params.Limit).Find(&product).Error
	if err != nil {
		tx.Rollback()
	}

	return &product, err
}

func (r *productRepository) CategoryExist(id uuid.UUID) error {
	var productCategory dto.ProductCategory

	err := r.db.Model(&entity.ProductCategory{}).Select("id").First(&productCategory, "id = ?", id).Error
	
	return err
}

func (r *productRepository) Create(product *entity.Product) error {
	err := r.db.Create(&product).Error

	return err
}

