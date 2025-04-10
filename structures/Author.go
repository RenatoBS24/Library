package structures

import (
	"Library/tools"
	"fmt"
	"strconv"
	"strings"
)

type Author struct {
	Name     string
	LastName string
	Age      int
}

func NewAuthor(name string, lastname string, age int) *Author {
	return &Author{name, lastname, age}
}
func CreateAuthors() []*Author {
	authors := make([]*Author, 0)
	data := tools.ReadFile("resources/authors.txt")
	for i := 0; i < len(data); i++ {
		dataDiv := strings.Split(data[i], ",")
		age, err := strconv.Atoi(dataDiv[2])
		if err != nil {
			fmt.Println("Error al convertir el tipo de dato:", err)
		}
		authors = append(authors, NewAuthor(dataDiv[0], dataDiv[1], age))
	}
	return authors
}
