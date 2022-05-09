package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

type AppConfig struct {
	Port     string
	Database struct {
		Address string
		Name    string
		Port    string
	}
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	defaultConfig.Port = os.Getenv("APP_PORT")
	defaultConfig.Database.Name = os.Getenv("DB_NAME")
	defaultConfig.Database.Address = os.Getenv("DB_ADDRESS")
	defaultConfig.Database.Port = os.Getenv("DB_PORT")

	return &defaultConfig
}
