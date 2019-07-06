package finder

import (
	"encoding/csv"
	"fmt"
	"github.com/hikaru7719/bookshelf-bot/domain"
	"io"
	"os"
	"strings"
)

var bookPath = "/Users/hikaru/develop/bookshelf-bot/book.csv"

type Finder interface {
	Find(searchWord string) ([]domain.Book, error)
	Close()
}

func NewCSV() (Finder, error) {
	file, err := os.Open(bookPath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("File Open OK.")
	return &CSV{reader: file}, nil
}

type CSV struct {
	reader io.ReadCloser
}

func (c *CSV) Find(searchWord string) ([]domain.Book, error) {
	bookSlice := make([]domain.Book, 0, 10)
	csvReader := csv.NewReader(c.reader)
	record, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for _, row := range record {
		if strings.Contains(row[2], searchWord) && row[2] != "" && searchWord != "" {
			newBook := domain.Book{ISBN: row[0], Title: row[2], Author: row[3], Publisher: row[4]}
			bookSlice = append(bookSlice, newBook)
		}
	}
	return bookSlice, nil
}

func (c *CSV) Close() {
	c.reader.Close()
}
