package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Product 1", 100)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "Product 1", p.Name)
	assert.Equal(t, 100.0, p.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 100.0)
	assert.NotNil(t, p)
	assert.Equal(t, ErrNameIsRequiered, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("Product 1", 0)
	assert.NotNil(t, p)
	assert.Equal(t, ErrPriceIsRequiered, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct("Product 1", -10)
	assert.NotNil(t, p)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestProductValidate(t *testing.T) {
	p, err := NewProduct("Product 1", 100)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}