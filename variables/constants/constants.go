package main

/**
 * Constants
 *
 * Constants are unique in that they only exist at compile time. They are
 * never allocated on the stack, nor on the heap. They live in the data
 * section of the program, and are write protected.
 *
 * Constants provide 256 bits of precision. This means that we can store
 * data which is much larger than you can easily represent in Golang, as
 * the native types only go up to 64 bits.
 */

// Duration represents time.
type Duration int64

const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
	Day                  = 24 * Hour
	Week                 = 7 * Day
)

// iota is a way to provide an atomically incrementing value to
// each constant within a block. This makes it much easier to
// create a set of constants (which represent some concept)
const (
	A1 = iota
	B1 = iota
	C1 = iota
)

const (
	A2 = iota // 0
	B2        // 1
	C2        // 2
)

const (
	A3 = iota + 1 // 1
	B3            // 2
	C3            // 3
)

const (
	V1 = 1 << iota // 1
	V2             // 2
	V3             // 4
	V4             // 8
	V5             // 16
)

func main() {
}
