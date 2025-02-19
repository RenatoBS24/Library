package Structures

import "fmt"

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
