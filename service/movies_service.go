package service

import (
	"context"
	"log"
	"sync"

	"github.com/ThailanTec/challenger/movies/domain"
	"github.com/ThailanTec/challenger/movies/infra/database/repositories"
)

type MovieServices interface {
	SaveMovie(data *domain.Movie) error
	GetMoviesDB() ([]*domain.Movie, error)
	Work(ctx context.Context, wg *sync.WaitGroup, movie <-chan *domain.Movie)
}

type movieService struct {
	movieRepo     repositories.MovieRepository
	movieSQLXrepo repositories.MovieSQLXRepository
}

func NewMovieService(mr repositories.MovieRepository, mrsqlx repositories.MovieSQLXRepository) movieService {
	return movieService{
		movieRepo:     mr,
		movieSQLXrepo: mrsqlx,
	}
}
func (ms *movieService) SaveMovies(data []*domain.Movie) error {
	err := ms.movieSQLXrepo.CreateMovies(data)
	if err != nil {
		return err
	}

	return nil
}

func (ms *movieService) GetMoviesDB() ([]*domain.Movie, error) {
	users, err := ms.movieRepo.GetMovies()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ms *movieService) Work(ctx context.Context, wg *sync.WaitGroup, moviesChan <-chan []*domain.Movie) {
	defer wg.Done()

	for {
		select {
		case moviesBatch, ok := <-moviesChan:
			if !ok {
				log.Println("Canal fechado, encerrando goroutine")
				return
			}

			err := ms.movieSQLXrepo.CreateMovies(moviesBatch) // Salva o batch de filmes
			if err != nil {
				log.Printf("Erro ao salvar o batch de filmes: %v", err)
				continue
			}

			log.Printf("Batch de %d filmes salvo com sucesso", len(moviesBatch))

		case <-ctx.Done():
			log.Println("Contexto cancelado, encerrando goroutine")
			return
		}
	}
}
