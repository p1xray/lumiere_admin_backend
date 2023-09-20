package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Конфиг
type Config struct {
	// Адрес, на котором запускается сервер
	HttpAddr string
	// Строка подключения к БД
	DatabaseUrl string

	// Количество пул соединений
	PoolMax int
}

// Читает конфиг из переменных окружения
func Read() Config {
	var config Config

	if err := godotenv.Load("local.env"); err != nil {
		log.Fatalf("Config Read Error: %v", err)
	}

	httpAddr, exists := os.LookupEnv("HTTP_ADDR")
	if exists {
		config.HttpAddr = httpAddr
	}

	dbUrl, exists := os.LookupEnv("DB_URL")
	if exists {
		config.DatabaseUrl = dbUrl
	}

	poolMaxStr, exists := os.LookupEnv("POOL_MAX")
	if exists {
		poolMax, err := strconv.Atoi(poolMaxStr)
		if err == nil {
			config.PoolMax = poolMax
		}
	}

	return config
}
