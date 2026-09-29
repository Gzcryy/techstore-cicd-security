package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupProductRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	router := gin.New()

	router.GET("/products", getProducts)
	router.GET("/products/:id", getProduct)
	router.POST("/products", createProduct)
	router.PUT("/products/:id", updateProduct)
	router.DELETE("/products/:id", deleteProduct)

	return router
}

func TestGetProductNotFound(t *testing.T) {
	router := setupProductRouter()

	req := httptest.NewRequest(http.MethodGet, "/products/999", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", rec.Code)
	}
}

func TestCreateProductInvalidJSON(t *testing.T) {
	router := setupProductRouter()

	body := strings.NewReader(`{"name":`)
	req := httptest.NewRequest(http.MethodPost, "/products", body)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", rec.Code)
	}
}

func TestCreateProductInvalidData(t *testing.T) {
	router := setupProductRouter()

	body := strings.NewReader(`{
		"name": "",
		"description": "Produto inválido",
		"price": -100,
		"stock": -1
	}`)

	req := httptest.NewRequest(http.MethodPost, "/products", body)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", rec.Code)
	}
}

func TestUpdateProductNotFound(t *testing.T) {
	router := setupProductRouter()

	body := strings.NewReader(`{
		"name": "Produto",
		"description": "Produto teste",
		"price": 100,
		"stock": 10
	}`)

	req := httptest.NewRequest(http.MethodPut, "/products/999", body)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", rec.Code)
	}
}

func TestDeleteProductNotFound(t *testing.T) {
	router := setupProductRouter()

	req := httptest.NewRequest(http.MethodDelete, "/products/999", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", rec.Code)
	}
}
