package repositories

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ThailanTec/challenger/movies/domain"
)

type CSVReader interface {
	ReadRecords(fileName string) ([]*domain.Movie, error)
}

type csvReader struct{}

func NewCSVReader() CSVReader {
	return &csvReader{}
}

func (r *csvReader) ReadRecords(fileName string) (movies []*domain.Movie, err error) {
	name := fileName
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1
	_, err = reader.Read()
	if err != nil {
		return nil, fmt.Errorf("erro ao ignorar a primeira linha: %v", err)
	}

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var result []*domain.Movie
	for _, record := range records {
		result = append(result, &domain.Movie{
			MovieID: toInt(record[0]),
			Title:   record[1],
			Year:    toSplit(record[1]),
			Genres:  record[2],
		})
	}
	return result, nil
}

func toSplit(s string) string {
	i := strings.Index(s, "(")
	j := strings.Index(s, ")")

	if i >= 0 && j > i {
		return s[i+1 : j]
	}

	return ""
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("Erro ao converter '%s' para inteiro: %v\n", s, err)
		return 0 // Retorna um valor padrão, como 0, em caso de erro
	}

	return i
}
