package main

/**
 * Escape Analysis
 */

type user struct {
	name  string
	email string
}

func main() {
	u1 := createUser1()
	u2 := createUser2()

	println("u1", &u1, "u2", u2)
}

// createUser1 uses value semantics - It creates a user, then returns
// a copy of that created value
//go:noinline
func createUser1() user {
	u := user{
		name:  "Sean",
		email: "sean.heintz@example.com",
	}
	println("createUser1", &u)

	return u
}

// createUser2 uses pointer semantics, returning a pointer to the object
// that it created. When the PC exits this function, however, we might expect
// the frame for this function to be cleaned up and wiped, thus corrupting
// the `*user` pointer. Golang performs static code analysis (specifically
// escape analysis) to determine if a created value should be created on
// the stack (which is ephemeral and self-cleaning) or if it should be allocated
// on the heap. This is especially interesting because we continue to work
// with `u` as though it had value semantics, but since the allocation actually
// lives on the heap, golang will actually treat the allocated value as a pointer
// and rewrite the accesses as though they were via pointer semantics.
//
// Data allocated on the heap requires garbage collection.
//go:noinline
func createUser2() *user {
	u := user{
		name:  "Sean",
		email: "sean.heintz@example.com",
	}
	println("createUser2", &u)

	return &u
}
