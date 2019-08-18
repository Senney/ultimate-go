package main

/**
	Pointers are used to share data between function frames.
**/

func main() {
	count := 10

	increment(&count)
	println("count:\tValue of [", count, "]\tAddr of [", &count, "]")
}

func increment(inc *int) {
	*inc++
	println("inc:\tValue of [", *inc, "]\tAddr of [", inc, "]")
}
