package main

import (
	"cinema_v1/internal/handlers"
	"cinema_v1/internal/repository"
	"cinema_v1/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация хранилища в памяти
	movieRepo := repository.NewMovieMemoryRepository()

	// Инициализация сервиса
	movieService := service.NewMovieService(movieRepo)

	// Инициализация обработчиков
	movieHandler := handlers.NewMovieHandler(movieService)

	// Создание роутера Gin
	router := gin.Default()

	// Маршруты для фильмов
	router.GET("/movies", movieHandler.GetAllMovies)
	router.GET("/movies/:id", movieHandler.GetMovie)
	router.POST("/movies", movieHandler.CreateMovie)
	router.PUT("/movies/:id", movieHandler.UpdateMovie)
	router.DELETE("/movies/:id", movieHandler.DeleteMovie)

	// Запуск сервера
	router.Run(":8080")
}

// Get
// http://localhost:8080/movies
