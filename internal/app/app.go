package app

import (
	"api/test/catalog/internal/config"
	"api/test/catalog/internal/handler"
	"api/test/catalog/internal/repository"
	"api/test/catalog/internal/service"
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Run(cfg *config.Config) {
	dgpool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Connection to DB failed : %v", err)
	}
	defer dgpool.Close()

	productRepository := repository.NewPostgresProductRepository(dgpool)
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
