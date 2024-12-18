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
	err := config.LoadSettings(".env")
	if err != nil {
		log.Fatal(err)
		return
	}

	settings := config.GetSettings()

	db, err := internal.NewConnection(settings.FileDB["file_ip"], settings.FileDB["file_user"],
		settings.FileDB["file_password"], settings.FileDB["file_database"])
	if err != nil {
		log.Fatal(err)
		return
	}
	defer internal.Close(db)

	repo = generator.NewXlsxFileRepository(db)
	service := generator.NewGeneratorService(repo)
	handler := api.NewGeneratorHandler(service)
	router := api.NewGeneratorRouter(handler)
	router.Run(":8080")
}
