package main

import (
	"cinema_v1/internal/handlers"
	"cinema_v1/internal/repository"
	"cinema_v1/internal/service"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	// Инициализация хранилища в памяти
	repo := repository.NewMovieMemoryRepository()

	// Инициализация сервиса
	service := service.NewMovieService(repo)

	// Инициализация обработчиков
	handler := handlers.NewMovieHandler(service)

	// Запуск сервера
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler.InitRoutes(),
	}

	logrus.Info("Server starting on port 5050...")

	if err := server.ListenAndServe(); err != nil {
		logrus.Fatal(err)
	}
}

// Get
// http://localhost:8080/movies
