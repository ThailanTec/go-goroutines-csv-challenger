package repositories

import (
	"github.com/ThailanTec/challenger/movies/domain"
	"gorm.io/gorm"
)

type MovieRepository interface {
	CreateMovies(movie *domain.Movie) error
	GetMovies() ([]*domain.Movie, error)
}

type movieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) MovieRepository {
	return &movieRepository{db: db}
}

func (repo *movieRepository) CreateMovies(movie *domain.Movie) error {
	return repo.db.Create(movie).Error
}

func (repo *movieRepository) GetMovies() ([]*domain.Movie, error) {
	var users []*domain.Movie
	result := repo.db.Find(&users)
	return users, result.Error
}
