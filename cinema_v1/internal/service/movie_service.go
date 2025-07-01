package service

import (
	"cinema_v1/internal/models"
	"cinema_v1/internal/repository"
)

// Интерфейс сервиса (бизнес-логика)
type MovieService interface {
	GetAllMovies() ([]models.Movie, error)
	GetMovie(id int) (*models.Movie, error)
	CreateMovie(movie models.Movie) (*models.Movie, error)
	UpdateMovie(movie models.Movie) (*models.Movie, error)
	DeleteMovie(id int) error
}

// Реализация сервиса
type movieService struct {
	repo repository.MovieRepository
}

// Конструктор сервиса
func NewMovieService(repo repository.MovieRepository) MovieService {
	return &movieService{repo: repo}
}

// Получить все фильмы
func (s *movieService) GetAllMovies() ([]models.Movie, error) {
	return s.repo.GetAll()
}

// Получить фильм по ID
func (s *movieService) GetMovie(id int) (*models.Movie, error) {
	if id == 0 {
		return nil, nil
	}
	return s.repo.GetByID(id)
}

func (s *movieService) CreateMovie(movie models.Movie) (*models.Movie, error) {
	if movie.Title == "" {
		return nil, nil
	}
	return s.repo.Create(movie)
}

func (s *movieService) UpdateMovie(movie models.Movie) (*models.Movie, error) {
	if movie.ID == 0 {
		return nil, nil
	}
	return s.repo.Update(movie)
}

func (s *movieService) DeleteMovie(id int) error {
	if id == 0 {
		return nil
	}
	return s.repo.Delete(id)
}
