package main

import (
	"fmt"
)

// User .
type User interface {
	Login()
	Logout()
}

// Login .
func Login(u User) {
	u.Login()
}

// Logout .
func Logout(u User) {
	u.Logout()
}

// Student .
type Student struct{}

// Login .
func (s Student) Login() {
	fmt.Println("student login")
}

// Logout .
func (s Student) Logout() {
	fmt.Println("student logout")
}

// Teacher .
type Teacher struct{}

// Login .
func (t Teacher) Login() {
	fmt.Println("teacher login")
}

// Logout .
func (t Teacher) Logout() {
	fmt.Println("teacher logout")
}

func main() {
	t := Teacher{}
	s := Student{}

	users := []User{t, s}

	for _, user := range users {
		Login(user)
		Logout(user)
	}
}
