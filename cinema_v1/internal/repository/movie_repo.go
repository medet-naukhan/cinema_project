package repository

import (
	"cinema_v1/internal/models"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

// Интерфейс хранилища (абстракция для любой БД)
type MovieRepository interface {
	GetAll() ([]models.Movie, error)
	GetByID(id string) (*models.Movie, error)
	Create(movie models.Movie) (*models.Movie, error)
	Update(movie models.Movie) (*models.Movie, error)
	Delete(id string) error
}

// Реализация хранилища в БД
type MoviePostgresRepository struct {
	db *sql.DB
}

// Конструктор хранилища
func NewMoviePostgresRepository(db *sql.DB) MovieRepository {
	return &MoviePostgresRepository{db: db}
}

func (r *MoviePostgresRepository) EnsureTableExists() error {
	query := `
	CREATE TABLE IF NOT EXISTS movies (
	id UUID PRIMARY KEY
	title TEXT NOT NULL
	description TEXT 
	duration BIGINT
	rating REAL
	);`
	_, err := r.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create movies table: %w, err")
	}
	return nil
}

// Получить все фильмы
func (r *MoviePostgresRepository) GetAll() ([]models.Movie, error) {
	query := `SELECT id, title, description, duration, rating, genre FROM movies ORDER BY title ASC`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("falied to get all movies: %w", err)
	}
	defer rows.Close()

	movies := []models.Movie{}
	for rows.Next() {
		var movie models.Movie
		var durationNs int64

		if err := rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.Rating, &movie.Genre); err != nil {
			return nil, fmt.Errorf("failed to scan movie row: %w", err)
		}
		movie.Duration = time.Duration(durationNs)
		movies = append(movies, movie)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}

	return movies, nil
}

// Получить фильм по ID
func (r *MoviePostgresRepository) GetByID(id string) (*models.Movie, error) {
	var movie models.Movie
	var durationNs int64

	query := `SELECT id, title, description, duration, rating, genre FROM movies WHERE id = $1`

	row := r.db.QueryRow(query, id)

	err := row.Scan(&movie.ID, &movie.Title, &movie.Description, &durationNs, &movie.Rating, &movie.Genre)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get movie by ID: %w", err)
	}

	movie.Duration = time.Duration(durationNs)

	return &movie, nil
}

// Создать фильм
func (r *MoviePostgresRepository) Create(movie models.Movie) (*models.Movie, error) {
	// if movie.ID == "" {
	// 	newUUID, err := uuid.NewRandom()
	// 	if err != nil {
	// 		return nil, fmt.Errorf("failed to generate UUID: %w", err)
	// 	}
	// 	movie.ID = newUUID.String()
	// }

	query := `INSERT INTO movies (id, title, description, duration, rating, genre)
              VALUES ($1, $2, $3, $4, $5, $6)
              RETURNING id`

	var insertedID string

	err := r.db.QueryRow(query, movie.ID, movie.Title, movie.Description,
		movie.Duration.Nanoseconds(), movie.Rating, movie.Genre).Scan(&insertedID)
	if err != nil {
		return nil, fmt.Errorf("failed to create movie: %w", err)
	}

	movie.ID = insertedID
	return &movie, nil
}

// Обновить фильм
func (r *MoviePostgresRepository) Update(movie models.Movie) (*models.Movie, error) {
	query := `UPDATE movies SET title=$2, description=$3, duration=$4, rating=$5, genre=$6
              WHERE id=$1
              RETURNING id, title, description, duration, rating, genre`

	var updatedMovie models.Movie
	var durationNs int64

	err := r.db.QueryRow(query, movie.ID, movie.Title, movie.Description,
		movie.Duration.Nanoseconds(), movie.Rating, movie.Genre).
		Scan(&updatedMovie.ID, &updatedMovie.Title, &updatedMovie.Description,
			&durationNs, &updatedMovie.Rating, &updatedMovie.Genre)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to update movie: %w", err)
	}

	updatedMovie.Duration = time.Duration(durationNs)

	return &updatedMovie, nil
}

// Удалить фильм
func (r *MoviePostgresRepository) Delete(id string) error {
	query := `DELETE FROM movies WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete movie: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected after delete: %w", err)
	}

	if rowsAffected == 0 {
		// return errors.New("movie not found for deletion")
	}

	return nil
}
