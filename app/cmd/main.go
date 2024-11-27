package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	client "effective_project/app/internal/client"
	songDelivery "effective_project/app/internal/song/delivery"
	songRepository "effective_project/app/internal/song/repository"
	songUseacse "effective_project/app/internal/song/usecase"
	"effective_project/app/models"
)

// @title Music info API
// @version 1.0
// @description This is an example of a server for managing songs as a test task.
// @host localhost:8080
// @BasePath /songs
func main() {
	logrus.SetLevel(logrus.DebugLevel)
	err := godotenv.Load("app/cmd/configs/.env")
	if err != nil {
		log.Printf("Warning: Could not load .env file")
	}

	serverHost := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")
	postgresUser := os.Getenv("POSTGRES_USERNAME")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresDB := os.Getenv("POSTGRES_DATABASE")
	clientServerUrl := os.Getenv("CLIENT_SERVER_URL")

	if serverHost == "" || serverPort == "" || postgresUser == "" || postgresPassword == "" || postgresHost == "" || postgresPort == "" || postgresDB == "" {
		postgresHost = "localhost"
		postgresUser = "db_pg"
		postgresPassword = "db_postgres"
		postgresDB = "effective_project"
		postgresPort = "5432"
		serverPort = "8080"
	}

	config := fmt.Sprintf("host=%s user=%s password=%s database=%s port=%s", postgresHost, postgresUser, postgresPassword, postgresDB, postgresPort)
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	err = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error
	if err != nil {
		log.Fatalf("Failed to create extension uuid-ossp: %v", err)
	}

	err = db.AutoMigrate(&models.Song{})
	if err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	err = db.Exec(`CREATE INDEX IF NOT EXISTS idx_id ON songs (id);`).Error
	if err != nil {
		log.Fatalf("Failed to create index: %v", err)
	}

	e := echo.New()

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	songDB := songRepository.New(db)
	if clientServerUrl == "" {
		clientServerUrl = "http://localhost:8081"
	}
	clientUC := client.NewSongClient(clientServerUrl)
	songUC := songUseacse.New(songDB, clientUC)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	songDelivery.NewDelivery(e, songUC)
	logrus.Info("Server has started on host: http://localhost:8080")
	serverAddress := serverHost + ":" + serverPort
	e.Logger.Fatal(e.Start(serverAddress))
}
