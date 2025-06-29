package app

import (
	"api/test/catalog/internal/config"
	"api/test/catalog/internal/domain"
	"api/test/catalog/internal/handler"
	"api/test/catalog/internal/repository"
	"api/test/catalog/internal/service"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Run(cfg *config.Config) {
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Connection to DB failed : %v", err)
	}

	db.AutoMigrate(&domain.Product{})

	productRepository := repository.NewGormProductRepository(db)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/api/product", productHandler.GetAllProducts)
	r.Post("/api/product", productHandler.SaveProduct)
	r.Get("/api/product/{id}", productHandler.GetById)
	r.Delete("/api/product/{id}", productHandler.DeleteProductById)
	r.Put("/api/product/{id}", productHandler.UpdateProduct)

	log.Printf("The server starts on port %s\n", cfg.HTTPport)
	http.ListenAndServe(cfg.HTTPport, r)
}
