package repositories

import (
	"fmt"

	"github.com/ThailanTec/challenger/movies/domain"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type MovieSQLXRepository interface {
	CreateMovies(movies []*domain.Movie) error
}

type movieSQLXRepository struct {
	db *sqlx.DB
}

func NewMovieSQLXRepository(db *sqlx.DB) *movieSQLXRepository {
	return &movieSQLXRepository{db: db}
}

func (s *movieSQLXRepository) CreateMovies(movies []*domain.Movie) error {
	query := "INSERT INTO movies(movie_id, title, year, genres) VALUES "
	values := []interface{}{}

	for i, movie := range movies {
		num := i * 4
		query += fmt.Sprintf("($%d, $%d, $%d, $%d),", num+1, num+2, num+3, num+4)
		values = append(values, movie.MovieID, movie.Title, movie.Year, movie.Genres)
	}

	// Remover a última vírgula e adicionar RETURNING ou ON CONFLICT se necessário
	query = query[:len(query)-1]

	_, err := s.db.Exec(query, values...)
	if err != nil {
		return err
	}

	return nil
}
