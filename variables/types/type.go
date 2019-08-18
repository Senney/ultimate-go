/**
 * Types
 * ========================================================
 * Basic unit of memory in the Golang model is a "Byte". However, bytes are useless unless you
 * can associate a "type" with them. For example, 00001010 could be 10 if intepreted as an
 * integer, but could be something entirely different if associated with a "Color" or "Operation".
 */

package main

import "fmt"

func main() {
	// Initialize variables to their ZERO values - Memory is automatically initialized to 0.
	// Automatically setting memory to 0 fixes a huge number of bugs that are present in unsafe
	// languages like C/C++. This only happens with "var".
	var a int     // ? bytes - Based on the computer architecture - 8 bytes on arm64 - same as WORD size.
	var b string  // 16 bytes - empty string - 2 WORD data structure
	var c float64 // 8 bytes
	var d bool    // 1 byte

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n", d, d)

	// Declare and initialize variables - Not a good idea to use this for
	// initializing variables with their ZERO value. (e.g. aa := 0)
	aa := 10
	bb := "hello"
	cc := 3.14159
	dd := true

	fmt.Printf("aa := 10 \t %T [%v]\n", aa, aa)
	fmt.Printf("bb := \"hello\" \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 3.14159 \t %T [%v]\n", cc, cc)
	fmt.Printf("dd := true \t %T [%v]\n", dd, dd)

	// Go does not have casting, but it does have conversion.
	//  Casting is problematic in that it relies on the programmer knowing with certainty that
	//  the memory being cast matches the type that it is being cast to. It's a fast operation,
	//  but it doesn't preserve memory integrity.
	aaa := int32(10)
	fmt.Printf("aaa := int32(10) %T [%v]\n", aaa, aaa)
}
