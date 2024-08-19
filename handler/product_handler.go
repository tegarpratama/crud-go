package handler

import (
	"crud-go/dto"
	"crud-go/errorhandler"
	"crud-go/helper"
	"crud-go/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *productHandler {
	return &productHandler{
		service: service,
	}
}

func (h *productHandler) GetProducts(ctx *gin.Context) {
	filter := helper.FilterParams(ctx)
	categoryId := ctx.Query("category_id")
	posts, paginate, err := h.service.FindAll(filter, categoryId)
	
	if err != nil {
		errorhandler.HandleError(ctx, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "List product's",
		Paginate:   paginate,
		Data:       posts,
	})

	ctx.JSON(http.StatusOK, res)
}

func (h *productHandler) CreateProduct(ctx *gin.Context) {
	fmt.Println("-------------", ctx)
	var product dto.ProductRequest

	if err := ctx.ShouldBindJSON(&product); err != nil {
		errorhandler.HandleError(ctx, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if err := h.service.CreateProduct(&product); err != nil {
		errorhandler.HandleError(ctx, err)
		return
	}


	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message: "Successfully created product",
	})

	ctx.JSON(http.StatusCreated, res)
}