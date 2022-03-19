package book

import (
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

func (r *BookRepository) UpdateBook(book *Book) error {
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

func (r *BookRepository) InsertSampleData(books []Book) {
	for _, book := range books {
		r.db.Where(Book{Name: book.Name}).FirstOrCreate(&book)
	}
}
