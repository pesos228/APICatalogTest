package main

import (
	"api/test/catalog/internal/app"
	"api/test/catalog/internal/config"
)

func main() {
	cfg := config.New()
	app.Run(cfg)
}
