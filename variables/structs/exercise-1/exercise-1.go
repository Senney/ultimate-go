package main

import "fmt"

type user struct {
	name  string
	email string
	age   int
}

func main() {
	user1 := user{
		name:  "Sean Heintz",
		email: "sean.heintz@example.com",
		age:   26,
	}

	fmt.Printf("%+v", user1)
}
