package handlers

import (
	"net/http"

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

func (h *MovieHandler) GetAllMovies(c *gin.Context) {
	movies, err := h.service.GetAllMovies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func (h *MovieHandler) GetMovie(c *gin.Context) {
	id := c.Param("id")

	movie, err := h.service.GetMovie(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if movie == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func (h *MovieHandler) CreateMovie(c *gin.Context) {
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdMovie, err := h.service.CreateMovie(movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdMovie)
}

func (h *MovieHandler) UpdateMovie(c *gin.Context) {
	id := c.Param("id")

	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie.ID = id
	updatedMovie, err := h.service.UpdateMovie(movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if updatedMovie == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	c.JSON(http.StatusOK, updatedMovie)
}

func (h *MovieHandler) DeleteMovie(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteMovie(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted"})
}
