package configs

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var config Config

type Config struct {
	App appConfig
}

func Init() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf("please consider environment variables: %s\n", err)
	}
	err = envconfig.Process("", &config)
	if err != nil {
		log.Fatalf("parse config error: %v", err)
	}
	return &config
}

func GetConfig() *Config {
	return &config
}
