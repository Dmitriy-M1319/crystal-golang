package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Settings struct {
	BaseDB  map[string]string
	FileDB  map[string]string
	Storage string
}

func GetSettings(env string) (*Settings, error) {
	err := godotenv.Load(env)
	if err != nil {
		return nil, err
	}

	baseDb := make(map[string]string)
	baseDb["base_ip"] = os.Getenv("BASE_POSTGRES_IP")
	baseDb["base_user"] = os.Getenv("BASE_POSTGRES_USER")
	baseDb["base_password"] = os.Getenv("BASE_POSTGRES_PASSWORD")
	baseDb["base_database"] = os.Getenv("BASE_POSTGRES_DATABASE")

	fileDb := make(map[string]string)
	fileDb["file_ip"] = os.Getenv("FILE_POSTGRES_IP")
	fileDb["file_user"] = os.Getenv("FILE_POSTGRES_USER")
	fileDb["file_password"] = os.Getenv("FILE_POSTGRES_PASSWORD")
	fileDb["file_database"] = os.Getenv("FILE_POSTGRES_DATABASE")
	return &Settings{BaseDB: baseDb, FileDB: fileDb, Storage: os.Getenv("FILE_STORAGE")}, nil
}