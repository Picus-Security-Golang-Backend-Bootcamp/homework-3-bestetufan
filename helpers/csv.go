package helpers

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/bestetufan/bookstore/domain/author"
	"github.com/bestetufan/bookstore/domain/book"
)

func ReadBookCSV(filename string) ([]book.Book, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = ';'
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var result []book.Book

	for _, line := range lines[1:] {
		pageCount, _ := strconv.Atoi(line[3])
		price, _ := strconv.ParseFloat(line[4], 2)
		stockCount, _ := strconv.Atoi(line[5])

		book := book.NewBook(
			line[0],                             // Name
			line[1],                             // StockCode
			line[2],                             // ISBN
			pageCount,                           // PageCount
			price,                               // Price
			stockCount,                          // StockCount
			*author.NewAuthor(line[6], line[7]), // Author
		)
		result = append(result, *book)
	}

	return result, nil
}
