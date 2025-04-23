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

func NewAuthor(name string, lastName string, age int) *Author {
	return &Author{name, lastName, age}
}

func CreateAuthors() []*Author {
	authors := make([]*Author, 0)
	data := tools.ReadFile("resources/Authors.txt")
	for i := 0; i < len(data); i++ {
		dataDiv := strings.Split(data[i], ",")
		age, err := strconv.Atoi(dataDiv[2])
		if err != nil {
			fmt.Print("Error al convertir el tipo de dato :", err)
		}
		authors = append(authors, NewAuthor(dataDiv[0], dataDiv[1], age))
	}
	return authors
}
func (author *Author) PrintAuthor() {
	fmt.Println("-------------------------")
	fmt.Print("Nombre: ", author.Name)
	fmt.Println(" ", author.LastName)
	fmt.Println("Edad: ", author.Age)
	fmt.Println("-------------------------")
}
