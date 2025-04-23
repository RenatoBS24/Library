package structures

import (
	"Library/tools"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct {
	Username      string
	password      string
	Email         string
	BooksBorrowed int
}

func NewUser(username string, password string, email string, booksBorrowed int) *User {
	return &User{username, password, email, booksBorrowed}
}

func (user *User) ValidUser(password, username string) bool {
	if password == user.password && username == user.Username {
		return true
	}
	return false
}
func (user *User) Reserve(book Book) {
	book.Reserve()
	user.BooksBorrowed++
}
func (user *User) Devolution(book Book) {
	book.Devolution()
	user.BooksBorrowed--
}

func (user *User) printUser() {
	fmt.Println("------------------------")
	fmt.Println("Username: ", user.Username)
	fmt.Println("Email: ", user.Email)
	fmt.Println("Books Borrowed: ", user.BooksBorrowed)
}

func CreateUsers() []*User {
	users := make([]*User, 0)
	fileInfo, err := os.Stat("resources/users.txt")
	if err != nil {
		fmt.Println("Error al leer el archivo: ", err)
	}
	if fileInfo.Size() == 0 {
		return users
	}
	data := tools.ReadFile("resources/users.txt")
	for i := 0; i < len(data); i++ {
		dataDiv := strings.Split(data[i], ",")
		booksBorrowed, err := strconv.Atoi(strings.TrimSpace(dataDiv[3]))
		if err != nil {
			fmt.Println("Error al convertir el tipo de dato:", err)
		}
		users = append(users, NewUser(dataDiv[0], dataDiv[1], dataDiv[2], booksBorrowed))
	}
	return users
}

func RegisterUser(user User) bool {
	file, err := os.OpenFile("resources/users.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error al crear el usuario", err)
		return false
	}
	text := user.Username + "," + user.password + "," + user.Email + "," + strconv.Itoa(user.BooksBorrowed) + "\n"
	num, err2 := file.WriteString(text)
	if err2 != nil {
		fmt.Println("Error al escribir en el archivo: ", err2)
		return false
	}
	if num == 0 {
		return false
	}
	err3 := file.Close()
	if err3 != nil {
		fmt.Println("Error al escribir el archivo: ", err3)
		return false
	}
	return true
}
