package main

import (
	"crud-go/config"
	"crud-go/dto"
	"crud-go/helper"
	"crud-go/router"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	config.LoadConfig()
	config.LoadDB()

	r := gin.Default()
	r.RedirectTrailingSlash = false

	api := r.Group("/api")
	api.GET("/check-health", func(ctx *gin.Context) {
		res := helper.Response(dto.ResponseParams{
			StatusCode: http.StatusOK,
			Message:    "It's work",
		})

		ctx.JSON(http.StatusOK, res)
	})

	router.PostRouter(api)
	
	return r
}

func TestCheckHealth(t *testing.T) {
	r := setupRouter()

	t.Run("it should return status code 200", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/api/check-health", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("it should match expected JSON", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/api/check-health", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		expectedBody := `{"code":200, "message":"It's work", "status":"success"}`
		assert.JSONEq(t, expectedBody, w.Body.String())
	})
}

func TestCreateProduct(t *testing.T) {
	r := setupRouter()

	t.Run("it should create product successfully", func(t *testing.T) {
		payload := `{
			"name": "Product Testing",
			"description": "Description Testing",
			"category_id": "9e09c5ad-5e2c-11ef-8e2b-9c6b005912a0"
		}`

		req, _ := http.NewRequest(http.MethodPost, "/api/products/", strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("it should cannot create product with empty field", func(t *testing.T) {
		payload := `{
			"description": "Description Testing",
			"category_id": "9e09c5ad-5e2c-11ef-8e2b-9c6b005912a0"
		}`

		req, _ := http.NewRequest(http.MethodPost, "/api/products/", strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("it should cannot create product with wrong category id", func(t *testing.T) {
		payload := `{
			"name": "Product Testing",
			"description": "Description Testing",
			"category_id": "11111111-1111-1111-1111-9c6b005912a0"
		}`

		req, _ := http.NewRequest(http.MethodPost, "/api/products/", strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}
