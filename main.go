package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bestetufan/bookstore/domain/author"
	"github.com/bestetufan/bookstore/domain/book"
	"github.com/bestetufan/bookstore/helpers"
	"github.com/bestetufan/bookstore/infrastructure"
)

var (
	authorRepo *author.AuthorRepository
	bookRepo   *book.BookRepository
)

func init() {
	db := infrastructure.NewPostgresDB("host=localhost user=postgres password=postgres dbname=bookstore port=5432 sslmode=disable")
	authorRepo = author.NewAuthorRepository(db)
	bookRepo = book.NewBookRepository(db)

	authorRepo.Migration()
	bookRepo.Migration()

	bookRepo.InsertSampleData("book-data.csv")
}

func listBooks() {
	fmt.Println("List of books:")

	// Get all books from repo
	books, _ := bookRepo.GetAllBooks()

	// Display
	for _, book := range books {
		fmt.Println(book.ToString())
	}
}

func searchBook(query string) {
	// Find all books by query
	books := bookRepo.FindBooksByQuery(query)

	// Display
	for _, book := range books {
		fmt.Println(book.ToString())
	}
}

func buyBook(bookId int, count int) {
	// Check if count is greater than 0
	if count <= 0 {
		fmt.Println("Transaction count must be greater than zero!")
		return
	}

	// Get book from repo by ID
	book, err := bookRepo.GetBookById(bookId)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Check stock information
	if book.StockCount < count {
		fmt.Println("Not enough stock!")
		return
	}

	// Perform buy
	err = bookRepo.UpdateBookStock(book)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Operation completed successfully!")
	}
}

func deleteBook(bookId int) {
	// Get book from repo by ID
	book, err := bookRepo.GetBookById(bookId)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Perform delete
	err = bookRepo.DeleteBook(book)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Operation completed successfully!")
}

func main() {
	args := os.Args
	lowerCaseArgs := helpers.ToLowerSlice(args)

	// Display welcome message in case of no command sent
	if len(lowerCaseArgs) == 1 {
		fmt.Println("Command List")
		fmt.Println("-----------------")
		fmt.Println("Search Operation: \"search {keyword}\" \n",
			"List Operation: \"list\" \n",
			"Buy Operation: \"buy {bookId, count}\" \n",
			"Delete Operation: \"delete {bookId}\"")
		fmt.Println("-----------------")
		return
	}

	// Command logic
	switch lowerCaseArgs[1] {
	case "search":
		if len(lowerCaseArgs) < 3 {
			fmt.Println("Enter a book name to search!")
		} else {
			searchBook(strings.Join(lowerCaseArgs[2:], " "))
		}
	case "list":
		listBooks()
	case "buy":
		if len(lowerCaseArgs) != 4 {
			fmt.Println("Enter a book id and amount!")
		} else {
			// Convert and check parameters for type int
			bookId, err := strconv.Atoi(lowerCaseArgs[2])
			count, err := strconv.Atoi(lowerCaseArgs[3])

			if err != nil {
				fmt.Println("Parameters must be in correct type!")
			} else {
				buyBook(bookId, count)
			}
		}
	case "delete":
		if len(lowerCaseArgs) != 3 {
			fmt.Println("Enter a book id to delete!")
		} else {
			//Convert and check parameters for type int
			bookId, err := strconv.Atoi(lowerCaseArgs[2])

			if err != nil {
				fmt.Println("Parameters must be in correct type!")
			} else {
				deleteBook(bookId)
			}
		}
	default:
		fmt.Println("Unknown command!")
	}
}
