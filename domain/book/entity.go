package book

import (
	"fmt"

	"github.com/bestetufan/bookstore/domain/author"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name       string
	StockCode  string
	ISBN       string
	PageCount  int
	Price      float64
	StockCount int
	AuthorID   uint
	Author     author.Author `gorm:"foreignKey:AuthorID"`
}

// Constructor
func NewBook(name string, stockCode string, isbn string, pageCount int, price float64,
	stockCount int, author author.Author) *Book {
	book := &Book{
		Name:       name,
		StockCode:  stockCode,
		ISBN:       isbn,
		PageCount:  pageCount,
		Price:      price,
		StockCount: stockCount,
		Author:     author,
	}
	return book
}

func (Book) TableName() string {
	return "book"
}

func (b *Book) ToString() string {
	return fmt.Sprintf("ID: %d => Name: %s, Author: %s, Pages: %d, Stock Count: %d, ISBN: %s, Stock Code: %s, CreatedAt : %s]",
		b.ID, b.Name, b.Author.GetFullName(), b.PageCount, b.StockCount, b.ISBN, b.StockCode, b.CreatedAt.Format("2006-01-02 15:04:05"))
}
