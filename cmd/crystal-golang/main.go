package main

import (
	"log"

	"github.com/Dmitriy-M1319/crystal-golang/api"
	"github.com/Dmitriy-M1319/crystal-golang/config"
	"github.com/Dmitriy-M1319/crystal-golang/internal"
	baseapp "github.com/Dmitriy-M1319/crystal-golang/internal/base-app"
	"github.com/Dmitriy-M1319/crystal-golang/internal/generator"
)

var fRepo generator.IFileRepository
var uRepo baseapp.IUserRepository
var oRepo baseapp.IOrderRepository
var pRepo baseapp.IProductRepository

func main() {
	err := config.LoadSettings(".env")
	if err != nil {
		log.Fatal(err)
		return
	}

	settings := config.GetSettings()

	fileDb, err := internal.NewConnection(settings.FileDB["ip"], settings.FileDB["port"], settings.FileDB["user"],
		settings.FileDB["password"], settings.FileDB["database"])
	if err != nil {
		log.Fatal(err)
		return
	}
	defer internal.Close(fileDb)
	baseDb, err := internal.NewConnection(settings.BaseDB["ip"], settings.BaseDB["port"], settings.BaseDB["user"],
		settings.BaseDB["password"], settings.BaseDB["database"])
	if err != nil {
		log.Fatal(err)
		return
	}
	defer internal.Close(baseDb)

	uRepo = baseapp.NewSqlUserRepository(baseDb)
	oRepo = baseapp.NewSqlOrderRepository(baseDb)
	pRepo = baseapp.NewSqlProductRepository(baseDb)
	ordService := baseapp.NewOrderService(uRepo, oRepo, pRepo)

	fRepo = generator.NewXlsxFileRepository(fileDb)
	service := generator.NewGeneratorService(fRepo)
	handler := api.NewGeneratorHandler(service, ordService)
	router := api.NewGeneratorRouter(handler)
	router.Run(":8080")
}
