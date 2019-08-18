package main

import "fmt"

type bill struct {
	counter int
	flag    bool
	pi      float32
}

type alice struct {
	counter int
	flag    bool
	pi      float32
}

func main() {
	// Create an anonymous struct and initialize it to 0.
	var e1 struct {
		counter int
		flag    bool
		pi      float32
	}
	fmt.Printf("var e1 struct\t%+v\n", e1)

	// Create an anonymous struct and initialize it with a literal constructor
	e2 := struct {
		counter int
		flag    bool
		pi      float32
	}{
		counter: 10,
		flag:    true,
		pi:      3.14159,
	}

	var b bill
	var b2 bill
	var a alice
	// b = a // golang does not do implicit conversions of data.

	// Explicit conversion helps to avoid some of the classes of bug that come from
	// implicit conversion of data. This makes a copy of `a`, but as type `bill`
	b = bill(a)
	fmt.Printf("%+v %+v\n", b, a)

	a.pi = 3.14159
	fmt.Printf("%v %v\n", b.pi, a.pi)

	// With anonymous structs however, implicit conversion is permitted as the
	// type is not **named**.
	b2 = e2
	fmt.Printf("%+v\n", b2)
}
