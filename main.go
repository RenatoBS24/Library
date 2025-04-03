package main

import (
	"Library/Structures"
	"bufio"
	"fmt"
	"os"
	"time"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	Index()
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
	User1 := Structures.NewUser("Renato", "1234", "renato@gamil.com", 0)
	User2 := Structures.NewUser("Martin", "12345", "martin@gmail.com", 0)
	Users = append(Users, User1)
	Users = append(Users, User2)
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
		} else {
			fmt.Println("Usuario o contraseña incorrecta")
			Login(Users)
		}
	case option == "2":
		Register(Users)
	case option == "3":
		os.Exit(0)
	default:
		fmt.Println("Opción no válida")
	}
}
func Login(users []*Structures.User) bool {
	confirmation := false
	fmt.Println("***************************")
	fmt.Println("Iniciar Sesión")
	fmt.Println("***************************")
	fmt.Println("Username: ")
	scanner.Scan()
	username := scanner.Text()
	fmt.Println("Password: ")
	scanner.Scan()
	password := scanner.Text()
	for _, User := range users {
		if User.ValidUser(password, username) {
			confirmation = true
			break
		}
	}
	return confirmation
}

func Register(users []*Structures.User) bool {
	username := ""
	password := ""
	password2 := ""
	email := ""
	fmt.Println("***************************")
	fmt.Println("Registrarse")
	fmt.Println("***************************")
	fmt.Println("Ingrese su nombre de usuario: ")
	scanner.Scan()
	username = scanner.Text()
	fmt.Println("Ingrese su email: ")
	scanner.Scan()
	email = scanner.Text()
	fmt.Println("Ingrese su contraseña: ")
	scanner.Scan()
	password = scanner.Text()
	fmt.Println("Confirme su contraseña: ")
	scanner.Scan()
	password2 = scanner.Text()
	if password == password2 {
		user := Structures.NewUser(username, password, email, 0)
		users = append(users, user)
		return true
	} else {
		return false
	}
}
func Menu(books []*Structures.Book) {
	fmt.Println("***************************")
	fmt.Println("Menú")
	fmt.Println("***************************")
	fmt.Println("1. Ver libros")
	fmt.Println("2. Reservar libro")
	fmt.Println("3. Devolver libro")
	fmt.Println("4. Volver")
	scanner.Scan()
	option := scanner.Text()
	switch {
	case option == "1":
		for _, book := range books {
			book.PrintBook()
			fmt.Println("-------------------------")
		}
	case option == "2":
		bookName := " "
		fmt.Println("Ingrese el nombre del libro:")
		scanner.Scan()
		bookName = scanner.Text()
		fmt.Println("El libro seleccionado es: " + bookName)
		for _, Book := range books {
			if Book.Title == bookName {
				Book.Available = false
			}
		}
	case option == "3":
		bookName := ""
		fmt.Println("Ingrese el nombre del libro:")
		scanner.Scan()
		bookName = scanner.Text()
		for _, Book := range books {
			if Book.Title == bookName {
				Book.Available = true
			}
		}
		fmt.Println("Libro devuelto correctamente")
	case option == "4":
		Index()
	default:
		fmt.Println("Opción no válida")
	}

}
