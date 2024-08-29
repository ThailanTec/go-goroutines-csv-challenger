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
	movieRepo repositories.MovieRepository
}

func NewMovieService(mr repositories.MovieRepository) movieService {
	return movieService{
		movieRepo: mr,
	}
}

func (ms *movieService) SaveMovie(data *domain.Movie) error {
	err := ms.movieRepo.CreateMovies(data)
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

func (ms *movieService) Work(ctx context.Context, wg *sync.WaitGroup, movie <-chan *domain.Movie) {
	defer wg.Done()

	for {
		select {
		case movies, ok := <-movie:
			if !ok {
				log.Println("Canal fechado, encerrando goroutine")
				return
			}

			err := ms.movieRepo.CreateMovies(movies)
			if err != nil {
				log.Printf("Erro ao subir o filme %v: %v", movies.MovieID, err)
				continue
			}

			log.Printf("Filme %v salvo com sucesso", movies.MovieID)

		case <-ctx.Done():
			log.Println("Contexto cancelado, encerrando goroutine")
			return
		}
	}

}
