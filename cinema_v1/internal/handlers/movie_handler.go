package handlers

import (
	"net/http"
	"strconv"

	"cinema_v1/internal/models"
	"cinema_v1/internal/service"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	service service.MovieService
}

func NewMovieHandler(service service.MovieService) *MovieHandler {
	return &MovieHandler{service: service}
}

func (h *MovieHandler) InitRoutes() *gin.Engine {
	router := gin.Default()
	// base middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//Регистрируем маршруты
	router.GET("/movies", h.GetAllMovies)
	router.GET("/movies/:id", h.GetMovie)
	router.POST("/movies", h.CreateMovie)
	router.PUT("movies/:id", h.UpdateMovie)
	router.DELETE("/movies/:id", h.DeleteMovie)

	return router
}

// Получить все фильмы
func (h *MovieHandler) GetAllMovies(c *gin.Context) {
	movies, err := h.service.GetAllMovies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movies)
}

// Получить фильм по ID
func (h *MovieHandler) GetMovie(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id param is not correct"})
		return
	}

	movie, err := h.service.GetMovie(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if movie == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	c.JSON(http.StatusOK, movie)
}

// Создать фильм
func (h *MovieHandler) CreateMovie(c *gin.Context) {
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	createdMovie, err := h.service.CreateMovie(movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if createdMovie == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie data"})
		return
	}

	c.JSON(http.StatusCreated, createdMovie)
}

// Обновить фильм
func (h *MovieHandler) UpdateMovie(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id param is not correct"})
		return
	}

	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	movie.ID = id
	updatedMovie, err := h.service.UpdateMovie(movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if updatedMovie == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	c.JSON(http.StatusOK, updatedMovie)
}

// Удалить фильм
func (h *MovieHandler) DeleteMovie(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id param is not correct"})
		return
	}

	err = h.service.DeleteMovie(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted"})
}
