package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
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
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	//r.Use(LoadRequest)
	//r.Use(middleware.RealIP)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("jwtExpiresIn", configs.JwtExpiresIn))
	
	//Products
	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.Create)
		r.Get("/", productHandler.GetAll)
		r.Get("/{id}", productHandler.Get)
		r.Put("/{id}", productHandler.Update)
		r.Delete("/{id}", productHandler.Delete)
	})
	

	//Users
	r.Post("/users", userHandler.Create)
	r.Post("/users/generate_token", userHandler.GetJWT)
	//r.Get("/users/{email}", userHandler.Get)
	//r.Put("/users/{email}", userHandler.Update)

	//http.HandleFunc("/products", productHandler.CreateProduct)

	http.ListenAndServe(":8000", r)
}

func LoadRequest(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		log.Println(r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
		next.ServeHTTP(w, r)
	})
}
