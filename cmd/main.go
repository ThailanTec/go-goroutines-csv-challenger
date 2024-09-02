package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/ThailanTec/challenger/movies/domain"
	"github.com/ThailanTec/challenger/movies/infra/database"
	"github.com/ThailanTec/challenger/movies/infra/database/repositories"
	"github.com/ThailanTec/challenger/movies/service"
	"github.com/ThailanTec/challenger/movies/src/config"
)

func main() {
	cnf := config.LoadConfig()
	db, err := database.PostgresClient(cnf)
	if err != nil {
		log.Fatal("Erro ao carregar o banco de dados", err)
	}

	sqlx, err := database.PostgresSQLXClient(cnf)
	if err != nil {
		log.Fatal("Erro ao carregar o banco de dados", err)
	}

	start := time.Now()

	CSVrepo := repositories.NewCSVReader()
	DBrepo := repositories.NewMovieRepository(db)
	dbSQLX := repositories.NewMovieSQLXRepository(sqlx)

	srv := service.NewMovieService(DBrepo, dbSQLX)

	movies, err := CSVrepo.ReadRecords("movie.csv")
	if err != nil {
		log.Fatal("Erro ao trazer os filmes patrão", err)
	}

	// go routines
	chunkSize := 100 // Define o tamanho do chunk
	ch := make(chan []*domain.Movie, chunkSize)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go srv.Work(context.Background(), &wg, ch)
	}

	go func() {
		chunks := chunkMovies(movies, chunkSize)
		for _, chunk := range chunks {
			ch <- chunk
			log.Printf("Chunk de %d filmes enviado", len(chunk))
		}
		close(ch)
	}()

	wg.Wait()
	elapsed := time.Since(start) // Finaliza o cronômetro

	log.Printf("Processo completo em %s", elapsed)
}

func chunkMovies(movies []*domain.Movie, chunkSize int) [][]*domain.Movie {
	var chunks [][]*domain.Movie
	for chunkSize < len(movies) {
		movies, chunks = movies[chunkSize:], append(chunks, movies[0:chunkSize:chunkSize])
	}

	chunks = append(chunks, movies)
	return chunks
}
