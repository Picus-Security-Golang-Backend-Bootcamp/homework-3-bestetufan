package book

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/bestetufan/bookstore/domain/author"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (r *BookRepository) GetAllBooks() ([]Book, error) {
	var books []Book
	result := r.db.Preload("Author").Find(&books)

	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

func (r *BookRepository) GetBookById(id int) (*Book, error) {
	var book *Book
	result := r.db.First(&book, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return book, nil
}

func (r *BookRepository) GetBookByName(name string) (*Book, error) {
	var book *Book
	result := r.db.Where(Book{Name: name}).Attrs(Book{}).FirstOrInit(&book)

	if result.Error != nil {
		return nil, result.Error
	}

	return book, nil
}

func (r *BookRepository) FindAllBooks() []Book {
	var books []Book
	r.db.Find(&books)

	return books
}

func (r *BookRepository) FindBooksByQuery(query string) []Book {
	var books []Book

	chain := r.db.Preload("Author").Where("name ILIKE ?", "%"+query+"%")
	chain = chain.Or("stock_code ILIKE ?", "%"+query+"%")
	chain = chain.Or("isbn ILIKE ?", "%"+query+"%")
	chain = chain.Find(&books)

	return books
}

func (r *BookRepository) UpdateBookStock(book *Book) error {
	// Update stock
	book.StockCount -= 1
	r.db.Save(&book)

	return nil
}

func (r *BookRepository) DeleteBook(book *Book) error {
	r.db.Delete(&book)

	return nil
}

func (r *BookRepository) Migration() {
	r.db.AutoMigrate(&Book{})
}

func (r *BookRepository) InsertSampleData(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = ';'
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, line := range lines[1:] {
		pageCount, _ := strconv.Atoi(line[3])
		price, _ := strconv.ParseFloat(line[4], 2)
		stockCount, _ := strconv.Atoi(line[5])

		book := NewBook(
			line[0],                             // Name
			line[1],                             // StockCode
			line[2],                             // ISBN
			pageCount,                           // PageCount
			price,                               // Price
			stockCount,                          // StockCount
			*author.NewAuthor(line[6], line[7]), // Author
		)
		r.db.Where(Book{Name: book.Name}).FirstOrCreate(&book)
	}
}
