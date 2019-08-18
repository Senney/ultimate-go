package main

/**
 * Stack Growth
 *
 * Stacks in golang are generally very small. Functions which make a ton of
 * calls could potentially run out of stack space given such a small stack
 * space. When we're out of stack, the Golang compiler will allocate another
 * larger stack (25% larger than the OG) and copy the existing stack in to
 * the newer stack.
 *
 * Surprisingly 2K is large enough for most operations in golang. The frames
 * tend to remain fairly small, especially after optimizations are applied by
 * the compiler.
 *
 * One of the constrains of Golang is that, since stacks can move, the only
 * pointer to a stack which is allowed is a local value. No pointers are
 * permitted to stacks across goroutines, for example.
 */

const size = 1024

func main() {
	s := "HELLO"
	stackCopy(&s, 0, [size]int{})
}

func stackCopy(s *string, c int, a [size]int) {
	println(c, s, *s)

	c++
	if c == 10 {
		return
	}

	stackCopy(s, c, a)
}

/**
$ go run lesson-4.go
0 0xc000063f78 HELLO
1 0xc000063f78 HELLO
<-- Stack copy happens here
2 0xc000073f78 HELLO
3 0xc000073f78 HELLO
4 0xc000073f78 HELLO
5 0xc000073f78 HELLO
<-- Stack copy happens here
6 0xc000093f78 HELLO
7 0xc000093f78 HELLO
8 0xc000093f78 HELLO
9 0xc000093f78 HELLO
**/
