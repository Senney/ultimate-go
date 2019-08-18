package main

/**
 * Garbage Collector
 *
 * Memory allocations in the heap are not cleaned up automatically
 * when the active frame changes. The garbage collector in Golang is
 * a concurrent process that tidies the stack.
 *
 * The garbage collector algorithm focuses on 3 core features
 *  - Smallest heap size - Keep resource usage low.
 *  - Stop the world time - Keep GC locks under 100 microseconds.
 *  - As a side effect of this, the GC can commandeer up to 25% of
 * 	  the available CPU capacity. On a quad-core machine, this could
 *    be an entire core.
 *
 * Garbage collector structure
 * - Write barrier is enabled - This is a stop-the-world event.
 *	- Every running goroutine must be halted, and the write barrier is
 *    is enabled.
 *  - The only "Safe" spot to stop a goroutine is when it makes a function
 *    call. Goroutines which run extremely tight loops will not be interrupted
 *    and may stall garbage collection.
 * - MARK heap allocations black if they're still in use
 * - SWEEP unused allocations
 */

func main() {

}
