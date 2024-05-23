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
	configs , err := configs.LoadConfig(".")
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

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, configs.TokenAuth, configs.JwtExpiresIn)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	//Products
	r.Post("/products", productHandler.Create)
	r.Get("/products/{id}", productHandler.Get)
	r.Put("/products/{id}", productHandler.Update)
	r.Delete("/products/{id}", productHandler.Delete)
	r.Get("/products", productHandler.GetAll)

	//Users
	r.Post("/users", userHandler.Create)
	r.Post("/users/generate_token", userHandler.GetJWT)
	//r.Get("/users/{email}", userHandler.Get)
	//r.Put("/users/{email}", userHandler.Update)

	//http.HandleFunc("/products", productHandler.CreateProduct)

	http.ListenAndServe(":8000", r)
}
