package Structures

import (
	"fmt"
	"time"
)

type Book struct {
	Title           string
	Author          Author
	Category        string
	Gender          string
	PublicationDate time.Time
	Available       bool
}

func NewBook(Title string, Author Author, Category string, Gender string, PublicationDate time.Time, Available bool) *Book {
	return &Book{Title, Author, Category, Gender, PublicationDate, Available}
}

func (book *Book) Reserve() {
	book.Available = false
}
func (book *Book) Devolution() {
	book.Available = true
}

func (book *Book) PrintBook() {
	fmt.Println("------------------------")
	fmt.Println("Title: ", book.Title)
	fmt.Println("Author: ", book.Author.Name, book.Author.LastName)
	fmt.Println("Category: ", book.Category)
	fmt.Println("Publication Date: ", book.PublicationDate.Format("02/01/2006"))
	if book.Available {
		fmt.Println("Available: Yes")
	} else {
		fmt.Println("Available: No")
	}
}
