package main

import "fmt"

func main() {
	user1 := struct {
		name  string
		email string
		age   int
	}{
		name:  "Sean Heintz",
		email: "sean.heintz@example.com",
		age:   26,
	}

	fmt.Printf("%+v\n", user1)
}
