package main

import (
	"cinema_v1/internal/handlers"
	"cinema_v1/internal/repository"
	"cinema_v1/internal/service"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	dbHost := "db"
	dbPort := "5432"
	dbUser := "user"
	dbPassword := "password"
	dbName := "moviesdb"

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logrus.Fatalf("Failed to open database connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		logrus.Fatalf("Failed to connect to database: %v", err)
	}
	logrus.Info("Successfully connected to PostgreSQL database!")

	defer func() {
		err := db.Close()
		if err != nil {
			logrus.Errorf("Error closing database connection: %v", err)
		}
		logrus.Info("Database connection closed.")
	}()

	// Инициализация хранилища в памяти
	repo := repository.NewMoviePostgresRepository(db)

	err = repo.(*repository.MoviePostgresRepository).EnsureTableExists()
	if err != nil {
		logrus.Fatalf("Failed to ensure database table exists: %v", err)
	}
	logrus.Info("Database table 'movies' checked/created successfully.")

	// Инициализация сервиса
	service := service.NewMovieService(repo)

	// Инициализация обработчиков
	handler := handlers.NewMovieHandler(service)

	// Настраиваем сервер
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler.InitRoutes(),
	}

	logrus.Info("Server starting on port 8080...")

	if err := server.ListenAndServe(); err != nil {
		logrus.Fatal(err)
	}
}
