package repository

import (
	"sync"

	"cinema_v1/internal/models"
)

// Интерфейс хранилища (абстракция для любой БД)
type MovieRepository interface {
	GetAll() ([]models.Movie, error)
	GetByID(id string) (*models.Movie, error)
	Create(movie models.Movie) (*models.Movie, error)
	Update(movie models.Movie) (*models.Movie, error)
	Delete(id string) error
}

// Реализация хранилища в памяти
type MovieMemoryRepository struct {
	movies map[string]models.Movie
	mu     sync.RWMutex
}

// Конструктор хранилища
func NewMovieMemoryRepository() MovieRepository {
	return &MovieMemoryRepository{
		movies: make(map[string]models.Movie),
	}
}

// Получить все фильмы
func (r *MovieMemoryRepository) GetAll() ([]models.Movie, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	movies := make([]models.Movie, 0, len(r.movies))
	for _, movie := range r.movies {
		movies = append(movies, movie)
	}
	return movies, nil
}

// Получить фильм по ID
func (r *MovieMemoryRepository) GetByID(id string) (*models.Movie, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	movie, exists := r.movies[id]
	if !exists {
		return nil, nil
	}
	return &movie, nil
}

// Создать фильм
func (r *MovieMemoryRepository) Create(movie models.Movie) (*models.Movie, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.movies[movie.ID] = movie
	return &movie, nil
}

// Обновить фильм
func (r *MovieMemoryRepository) Update(movie models.Movie) (*models.Movie, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.movies[movie.ID]; !exists {
		return nil, nil
	}

	r.movies[movie.ID] = movie
	return &movie, nil
}

// Удалить фильм
func (r *MovieMemoryRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.movies, id)
	return nil
}
