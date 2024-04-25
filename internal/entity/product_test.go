package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Product 1", 10)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, 10, product.Price)
}

func TestProductWhenNameIsRequired(t *testing.T){
	p, err := NewProduct("", 10)
	assert.Nil(t, p)
	assert.Equal(t, ErrNameRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T){
	p, err := NewProduct("Product 1", 0)
	assert.Nil(t, p)
	assert.Equal(t, ErrPriceRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T){
	p, err := NewProduct("Product 1", -10)
	assert.Nil(t, p)
	assert.Equal(t, ErrorInvalidPrice, err)
}

func TestProductValiodate(t *testing.T){
	p, err := NewProduct("Product 1", 10)
	assert.Nil(t, err)
	assert.NotNil(t,p)
	assert.Nil(t, p.Validate())
}