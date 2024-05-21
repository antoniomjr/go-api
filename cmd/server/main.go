package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/antoniomjr/go/9-apis/configs"
	"github.com/antoniomjr/go/9-apis/internal/entity"
	"github.com/antoniomjr/go/9-apis/internal/infra/database"
	"github.com/antoniomjr/go/9-apis/internal/infra/webserver/handlers"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)

	//http.HandleFunc("/products", productHandler.CreateProduct)

	http.ListenAndServe(":8000", r)
}
