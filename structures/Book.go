package structures

import (
	"Library/tools"
	"fmt"
	"os"
	"strings"
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
func CreateBooks() []*Book {
	authors := CreateAuthors()
	books := make([]*Book, 0)
	data := tools.ReadFile("resources/books.txt")
	for i := 0; i < len(data); i++ {
		dataDiv := strings.Split(data[i], ",")
		books = append(books, NewBook(dataDiv[0], *authors[i], dataDiv[1], dataDiv[2], time.Now(), true))
	}
	return books
}

func RegisterBook(book Book) bool {
	file, err := os.OpenFile("resources/books.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error al abrir el archivo: ", err)
		return false
	}
	text := book.Title + "," + book.Category + "," + book.Gender
	num, err := fmt.Fprint(file, text)
	if err != nil {
		fmt.Println("Error al escribir en el archivo: ", err)
		return false
	}
	if num == 0 {
		return false
	}
	err2 := file.Close()
	if err2 != nil {
		fmt.Println("Error al escribir el archivo: ", err)
		return false
	}
	return true
}
