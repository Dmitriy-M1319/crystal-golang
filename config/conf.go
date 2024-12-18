package config

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Settings struct {
	BaseDB  map[string]string
	FileDB  map[string]string
	Storage string
}

var (
	once     sync.Once
	settings *Settings
)

func LoadSettings(env string) error {
	var err error = nil
	once.Do(func() {
		err = godotenv.Load(env)
		if err != nil {
			return
		}

		baseDb := make(map[string]string)
		baseDb["ip"] = os.Getenv("BASE_POSTGRES_IP")
		baseDb["port"] = os.Getenv("BASE_POSTGRES_PORT")
		baseDb["user"] = os.Getenv("BASE_POSTGRES_USER")
		baseDb["password"] = os.Getenv("BASE_POSTGRES_PASSWORD")
		baseDb["database"] = os.Getenv("BASE_POSTGRES_DATABASE")

		fileDb := make(map[string]string)
		fileDb["ip"] = os.Getenv("FILE_POSTGRES_IP")
		fileDb["port"] = os.Getenv("FILE_POSTGRES_PORT")
		fileDb["user"] = os.Getenv("FILE_POSTGRES_USER")
		fileDb["password"] = os.Getenv("FILE_POSTGRES_PASSWORD")
		fileDb["database"] = os.Getenv("FILE_POSTGRES_DATABASE")
		settings = &Settings{BaseDB: baseDb, FileDB: fileDb, Storage: os.Getenv("FILE_STORAGE")}
	})

	return err
}

func GetSettings() *Settings {
	return settings
}
