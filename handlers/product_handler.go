package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/you/product-api/models"
	"github.com/you/product-api/store"
)

type ProductHandler struct {
	store *store.InMemoryStore
}

func NewProductHandler(s *store.InMemoryStore) *ProductHandler {
	return &ProductHandler{store: s}
}

// ListProducts godoc
// @Summary Get all products
// @Description Get list of all products
// @Tags products
// @Produce json
// @Success 200 {array} models.Product
// @Router /products [get]
func (h *ProductHandler) ListProducts(c *gin.Context) {
	products, err := h.store.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to list products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

// CreateProduct godoc
// @Summary Create new product
// @Description Adds a new product to the store
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.Product true "Product data"
// @Success 201 {object} models.Product
// @Failure 400 {object} map[string]string
// @Router /products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var p models.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, err := h.store.Create(&p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to create product"})
		return
	}
	c.JSON(http.StatusCreated, created)
}

// GetProduct godoc
// @Summary Get product by ID
// @Description Get a single product by its ID
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Product
// @Failure 404 {object} map[string]string
// @Router /products/{id} [get]
func (h *ProductHandler) GetProduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	p, err := h.store.Get(id)
	if err == store.ErrNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}
	c.JSON(http.StatusOK, p)
}

// UpdateProduct godoc
// @Summary Update existing product
// @Description Update product information by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body models.Product true "Updated product"
// @Success 200 {object} models.Product
// @Failure 404 {object} map[string]string
// @Router /products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var upd models.Product
	if err := c.ShouldBindJSON(&upd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	p, err := h.store.Update(id, &upd)
	if err == store.ErrNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}
	c.JSON(http.StatusOK, p)
}

// DeleteProduct godoc
// @Summary Delete product
// @Description Delete product by ID
// @Tags products
// @Param id path int true "Product ID"
// @Success 204 "No Content"
// @Failure 404 {object} map[string]string
// @Router /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.store.Delete(id); err == store.ErrNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
	} else {
		c.Status(http.StatusNoContent)
	}
}
