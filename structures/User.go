package structures

import (
	"Library/tools"
	"fmt"
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
	data := tools.ReadFile("resources/users.txt")
	for i := 0; i < len(data); i++ {
		dataDiv := strings.Split(data[i], ",")
		booksBorrowed, err := strconv.Atoi(dataDiv[3])
		if err != nil {
			fmt.Println("Error al convertir el tipo de dato:", err)
		}
		users = append(users, NewUser(dataDiv[0], dataDiv[1], dataDiv[2], booksBorrowed))
	}
	return users
}
