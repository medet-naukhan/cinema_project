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
