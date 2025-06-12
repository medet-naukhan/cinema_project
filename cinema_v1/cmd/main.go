package main

import (
	"cinema_v1/internal/handlers"
	"cinema_v1/internal/repository"
	"cinema_v1/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация хранилища в памяти
	repo := repository.NewMovieMemoryRepository()

	// Инициализация сервиса
	service := service.NewMovieService(repo)

	// Инициализация обработчиков
	handler := handlers.NewMovieHandler(service)

	// Настраиваем сервер
	router := gin.Default()

	//Регистрируем маршруты
	router.GET("/movies", handler.GetAllMovies)
	router.GET("/movies/:id", handler.GetMovie)
	router.POST("/movies", handler.CreateMovie)
	router.PUT("movies/:id", handler.UpdateMovie)
	router.DELETE("/movies/:id", handler.DeleteMovie)

	// Запускаем сервер
	router.Run(":8080")
}
