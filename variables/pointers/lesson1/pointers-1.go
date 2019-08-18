package main

import "fmt"

/**
	In tradition programming/processing structures, we typically deal with
	stacks, heaps and data segments. A traditional stack is allocated to be
	1 MB for a given thread to use. In Goroutines, only 2kb of memory is
	allocated in the initial stack.

	A goroutine only has access to the frame of memory from the stack that
	it has been granted access to. The frame is a layer of isolation and
	immutability such that the goroutine can only mutate data within its
	own memory space.

	Every time a function is called, a new frame is sliced off of the
	stack. Once the program counter is jumped in to the new function,
	the function can only create data within that frame. Parameters
	which are passed by value are copied in to this frame. The default
	in Golang is to "Pass by Value/Copy" instead of "Pass by Reference"
	like languages such as Javascript. Even complex objects like Structs
	are passed by value, which can be a costly operation when working with
	larger objects.
**/

type fancy struct {
	val int
}

func main() {
	count := 10
	println("count:\tValue Of [", count, "]\tAddr of [", &count, "]")

	increment(count)
	println("count:\tValue Of [", count, "]\tAddr of [", &count, "]")

	obj := fancy{
		val: 4,
	}

	incrementStruct(obj)
	fmt.Printf("obj:\t%+v\n", obj)
}

func increment(inc int) {
	inc++
	println("inc:\tValue of [", inc, "]\tAddr of [", &inc, "]")
}

func incrementStruct(obj fancy) {
	obj.val++
	fmt.Printf("obj:\t%+v\n", obj)
}
