package main

import (
	"log"
	"study-service/internal/config"
	"study-service/internal/logger"
	"study-service/internal/repository"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		// Logger'ı başlatmak için gereken konfigürasyon yüklenemediği için
		// standart log paketi ile hata basıp çıkıyoruz.
		log.Fatalf("Error loading config: %v", err)
	}

	logger := logger.New(cfg.Environment == "development")
	logger.Info().Msg("Config loaded successfully")

	dbPool, err := repository.NewConnection(cfg.DatabaseURL)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to database")
	}
	defer dbPool.Close()
	logger.Info().Msg("Successfully connected to the database")
}
