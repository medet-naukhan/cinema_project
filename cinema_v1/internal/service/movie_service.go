package service

import (
	"cinema_v1/internal/models"
	"cinema_v1/internal/repository"
)

type MovieService interface {
	GetAllMovies() ([]models.Movie, error)
	GetMovie(id string) (*models.Movie, error)
	CreateMovie(movie models.Movie) (*models.Movie, error)
	UpdateMovie(movie models.Movie) (*models.Movie, error)
	DeleteMovie(id string) error
}

type movieService struct {
	repo repository.MovieRepository
}

func NewMovieService(repo repository.MovieRepository) MovieService {
	return &movieService{repo: repo}
}

func (s *movieService) GetAllMovies() ([]models.Movie, error) {
	return s.repo.GetAll()
}

func (s *movieService) GetMovie(id string) (*models.Movie, error) {
	return s.repo.GetByID(id)
}

func (s *movieService) CreateMovie(movie models.Movie) (*models.Movie, error) {
	return s.repo.Create(movie)
}

func (s *movieService) UpdateMovie(movie models.Movie) (*models.Movie, error) {
	return s.repo.Update(movie)
}

func (s *movieService) DeleteMovie(id string) error {
	return s.repo.Delete(id)
}
