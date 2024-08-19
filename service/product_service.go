package service

import (
	"crud-go/dto"
	"crud-go/entity"
	"crud-go/errorhandler"
	"crud-go/helper"
	"crud-go/repository"
	"math"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ProductService interface{
	FindAll(params *dto.FilterParams, categoryId string) (*[]dto.ProductResponse, *dto.Paginate, error)
	CreateProduct(req *dto.ProductRequest) error
}

type productService struct {
	repository repository.ProductRepository
	validator  *validator.Validate
}

func NewProductService(r repository.ProductRepository) *productService {
	return &productService{
		repository: r,
		validator: validator.New(),
	}
}

func (s *productService) FindAll(params *dto.FilterParams, categoryId string) (*[]dto.ProductResponse, *dto.Paginate, error) {
	total, err := s.repository.TotalData(params, categoryId)
	if err != nil {
		return nil, nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	posts, err := s.repository.FindAll(params, categoryId)
	if err != nil {
		return nil, nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	paginate := &dto.Paginate{
		Total:      int(total),
		PerPage:    params.Limit,
		Page:       params.Page,
		TotalPages: int(math.Ceil(float64(total) / float64(params.Limit))),
	}

	return posts, paginate, nil
}


func (s *productService) CreateProduct(req *dto.ProductRequest) error {
	if err := s.validator.Struct(req); err != nil {
		return &errorhandler.BadRequestError{Message: err.Error()}
	}

	if err := s.repository.CategoryExist(req.CategoryID); err != nil {
		return &errorhandler.NotFoundError{Message: "Category not found"}
	}

	product := entity.Product{
		ID: uuid.New(),
		Name: helper.SanitizeInput(req.Name),
		Description: helper.SanitizeInput(req.Description),
		Category_id: req.CategoryID,
	}

	if err := s.repository.Create(&product); err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	return nil
}