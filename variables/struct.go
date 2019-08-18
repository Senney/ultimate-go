/**
 * Struct types
 * ======================================
 *
 */

package main

import (
	"fmt"
	"unsafe"
)

type example struct {
	flag bool
	// 1 bytes padding - int16 must start aligned on 2 byte boundary
	counter int16
	pi      float32
}

type example2 struct {
	flag bool
	// 3 bytes of padding - int32 must start aligned on 4 byte boundary
	counter int32
	pi      float32
}

func main() {
	// Create e1, set it to its zero value.
	var e1 example

	// How much memory is required to allocate `example`?
	//  1 bool, 2 int 16, 4 float32 = 7 bytes
	// But we also need to remember that alignment exists
	//  Ensure that values fit inside of the WORD boundaries of memory.
	//  Alignment pads out structs to fill the remaining space in a
	//  boundary.
	// On our 64 bit processor, we'll see padding added to our
	// struct to pad it out to 8 bytes. But where will it be padding be
	// added?
	//  __________________________________________________________
	//  |flag 1 |pad 1|  counter 2  |        pi  4                |
	//  ¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯
	// Optimizing too early to minimize padding is a bad idea. Structs should
	// start as readable as possible, but once a memory profile has indicated
	// that the struct is problematic, we can start to optimize.
	//
	// One good practice when optimizing the alignment of a struct is to
	// order the struct from largest to smallest in terms of attribute size.
	// However, the fields should generally be ordered in terms of correctness
	// rather than performance.
	fmt.Printf("var e1 example\t %T [%+v] [%v]\n", e1, e1, unsafe.Sizeof(e1))

	// Declare a variable of type example, and initialize. * Literal Construction *
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.14159,
	}

	fmt.Printf("var e1 example\t %T [%+v] [%v]\n", e2, e2, unsafe.Sizeof(e2))
}
