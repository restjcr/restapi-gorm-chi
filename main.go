package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/restjcr/restapi-gorm-chi/database"
	"github.com/restjcr/restapi-gorm-chi/handler"
	"github.com/restjcr/restapi-gorm-chi/model"
	"github.com/restjcr/restapi-gorm-chi/repository"
	"github.com/restjcr/restapi-gorm-chi/service"
	"log"
	"net/http"
)

func main() {
	db := database.NewConnection()
	err := db.AutoMigrate(&model.Product{})

	if err != nil {
		log.Fatal("Cannot automigrate product model", err.Error())
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Get("/products", productHandler.GetAllProductsHandler)
	r.Get("/products/{id}", productHandler.GetProductHandler)
	r.Post("/products", productHandler.CreateProductHandler)
	// TODO
	r.Put("/products/{id}", productHandler.UpdateProductHandler)
	r.Delete("/products/{id}", productHandler.DeleteProductHandler)

	http.ListenAndServe(":3000", r)
}
