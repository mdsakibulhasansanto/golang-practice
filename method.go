package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// Spelling mistake: NmaeGet -> NameGet
func (u User) NameGet() {
	fmt.Println(u.Name)
}

func (u *User) AgeGet() {
	u.Age++
	fmt.Println(u.Age)
}

/*
func main() {
	user := User{Name: "Santo", Age: 21}

	user.NameGet()
}

*/
