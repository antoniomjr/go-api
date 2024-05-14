package database

import (
	"fmt"
	"testing"
	"math/rand"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/antoniomjr/go/9-apis/internal/entity"
)


func TestCreateNewProduct(t *testing.T) {
	db := setupDatabase(t)

	product, err := entity.NewProduct("Product 1", 10.00)
	assert.NoError(t, err)
	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)
}



func TestFindAllProducts(t *testing.T) {
	db := setupDatabase(t)

	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(product)
	}

	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)
}

func TestFindProductById(t *testing.T) {
	db := setupDatabase(t)

	product, err := entity.NewProduct("Product 1", 10.00)
	assert.NoError(t, err)
	db.Create(product)

	productDB := NewProduct(db)
	productFound, err := productDB.FindById(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestUpdateProduct(t *testing.T) {
	db := setupDatabase(t)

	product, err := entity.NewProduct("Product 1", 10.00)
	assert.NoError(t, err)
	db.Create(product)

	productDB := NewProduct(db)
	productFound, err := productDB.FindById(product.ID.String())
	assert.NoError(t, err)
	productFound.Name = "Product 2"
	productFound.Price = 20.00
	err = productDB.Update(productFound)
	assert.NoError(t, err)

	productUpdated, err := productDB.FindById(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, productFound.Name, productUpdated.Name)
	assert.Equal(t, productFound.Price, productUpdated.Price)
}

func TestDeleteProduct(t *testing.T) {
	db := setupDatabase(t)

	product, err := entity.NewProduct("Product 1", 10.00)
	assert.NoError(t, err)
	db.Create(product)
	productDB := NewProduct(db)
	
	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)

	_, err = productDB.FindById(product.ID.String())
	assert.Error(t, err)
}


func setupDatabase(t *testing.T) *gorm.DB {
    db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
    if err != nil {
        t.Fatalf("Failed to open database: %v", err)
    }

    err = db.AutoMigrate(&entity.Product{})
    if err != nil {
        t.Fatalf("Failed to migrate database: %v", err)
    }

    return db
}