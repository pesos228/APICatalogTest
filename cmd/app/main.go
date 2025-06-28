package main

import (
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

func main() {
	databaseUrl := "postgres://postgres:vladick@localhost:5432/catalog?search_path=public"

	dgpool, err := pgxpool.New(context.Background(), databaseUrl)
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

	http.ListenAndServe(":3000", r)
}
