package main

import (
	"Library/Structures"
	"Time"
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	fmt.Println("Hello World")
}

func StartLibrary() ([]*Structures.Book, []*Structures.User) {
	Author1 := Structures.NewAuthor("Gabriel", "Garcia Marquez", 50)
	Author2 := Structures.NewAuthor("Julio", "Cortazar", 45)
	Author3 := Structures.NewAuthor("Mario", "Vargas Llosa", 60)
	// -------------------------------
	Book1 := Structures.NewBook("Cien años de soledad", *Author1, "Novela", "Ficción", time.Now(), true)
	Book2 := Structures.NewBook("Rayuela", *Author2, "Novela", "Ficción", time.Now(), true)
	Book3 := Structures.NewBook("La ciudad y los perros", *Author3, "Novela", "Ficción", time.Now(), true)
	Books := make([]*Structures.Book, 0)
	Books = append(Books, Book1)
	Books = append(Books, Book2)
	Books = append(Books, Book3)
	// ..................................
	Users := make([]*Structures.User, 0)
	User1 := Structures.NewUser("User1", "1234", "renato@gamil.com", 0)
	Users = append(Users, User1)
	return Books, Users
}

func Index() {
	Books := make([]*Structures.Book, 0)
	Users := make([]*Structures.User, 0)
	Books, Users = StartLibrary()
	fmt.Println("***************************")
	fmt.Println("Bienvenido a la biblioteca")
	fmt.Println("***************************")
	fmt.Println("1. Iniciar Sesión")
	fmt.Println("2. Registrarse")
	fmt.Println("3. Salir")
	scanner.Scan()
	option := scanner.Text()
	switch {
	case option == "1":
		if Login(Users) {
			Menu(Books)
		}
	case option == "2":
		// Register()
	case option == "3":
		os.Exit(0)
	default:
		fmt.Println("Opción no válida")
	}
}
func Login(users []*Structures.User) bool {
	fmt.Println("***************************")
	fmt.Println("Iniciar Sesión")
	fmt.Println("***************************")
	fmt.Println("Username: ")
	scanner.Scan()
	username := scanner.Text()
	fmt.Println("Password: ")
	password := scanner.Text()
	for _, User := range users {
		if User.ValidUser(password, username) {
			return true
		}

	}
	return false
}

func Menu(books []*Structures.Book) {
	fmt.Println("***************************")
	fmt.Println("Menú")
	fmt.Println("***************************")
	fmt.Println("1. Ver libros")
	fmt.Println("2. Reservar libro")
	fmt.Println("3. Devolver libro")
	fmt.Println("4. Salir")

}
