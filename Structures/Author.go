package Structures

type Author struct {
	Name     string
	LastName string
	Age      int
}

func NewAuthor(name string, lastname string, age int) *Author {
	return &Author{name, lastname, age}
}
