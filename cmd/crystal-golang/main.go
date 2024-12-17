package main

import (
	"github.com/Dmitriy-M1319/crystal-golang/api"
	"github.com/Dmitriy-M1319/crystal-golang/internal/generator"
)

func main() {
	service := generator.NewGeneratorService()
	handler := api.NewGeneratorHandler(service, ".env")
	router := api.NewGeneratorRouter(handler)
	router.Run(":8080")
}
