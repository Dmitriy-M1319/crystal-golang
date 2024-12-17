package main

import (
	"log"

	"github.com/Dmitriy-M1319/crystal-golang/api"
	"github.com/Dmitriy-M1319/crystal-golang/config"
	"github.com/Dmitriy-M1319/crystal-golang/internal"
	"github.com/Dmitriy-M1319/crystal-golang/internal/generator"
)

var repo generator.IFileRepository

func main() {
	settings, err := config.GetSettings(".env")
	if err != nil {
		log.Fatal(err)
		return
	}

	db, err := internal.NewConnection(settings.FileDB["file_ip"], settings.FileDB["file_user"],
		settings.FileDB["file_password"], settings.FileDB["file_database"])
	if err != nil {
		log.Fatal(err)
		return
	}
	defer internal.Close(db)

	repo = generator.NewXlsxFileRepository(db)
	service := generator.NewGeneratorService(repo)
	handler := api.NewGeneratorHandler(service, settings)
	router := api.NewGeneratorRouter(handler)
	router.Run(":8080")
}
