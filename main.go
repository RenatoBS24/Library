package main

import (
	"Library/structures"
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	Index()
}

func Index() {
	Books := structures.CreateBooks()
	Users := structures.CreateUsers()
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
		Register(Users)
	case option == "3":
		os.Exit(0)
	default:
		fmt.Println("Opción no válida")
	}
}
func Login(users []*structures.User) bool {
	confirmation := false
	for confirmation == false {
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
	}
	return confirmation
}

func Register(users []*structures.User) bool {
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
		user := structures.NewUser(username, password, email, 0)
		users = append(users, user)
		return true
	} else {
		return false
	}
}
func Menu(books []*structures.Book) {
	next := true
	for next {
		fmt.Println("***************************")
		fmt.Println("Menú")
		fmt.Println("***************************")
		fmt.Println("1. Ver libros")
		fmt.Println("2. Reservar libro")
		fmt.Println("3. Devolver libro")
		fmt.Println("4. Cerrar sesión")
		scanner.Scan()
		option := scanner.Text()
		switch {
		case option == "1":
			for _, book := range books {
				book.PrintBook()
				fmt.Println("-------------------------")
			}
		case option == "2":
			for _, book := range books {
				book.PrintBook()
				fmt.Println("-------------------------")
			}
			bookName := " "
			fmt.Println("Ingrese el nombre del libro:")
			scanner.Scan()
			bookName = scanner.Text()
			fmt.Println("El libro seleccionado es: " + bookName)
			for _, Book := range books {
				if Book.Title == bookName && Book.Available == true {
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
			next = false
			Index()
		default:
			fmt.Println("Opción no válida")
		}
	}
}
