# Notes

## The Difference Between Concurrency and Parallelism

* Concurrency is a property of the code; parallelism is a property of the running program.
    * The chunks of our program may appear to be running in parallel, but really they’re executing in a sequential manner faster than is distinguishable.
* We do not write parallel code, only concurrent code that we hope will be run in parallel.
* Layers of abstractions are what allow us to make the distinction between concurrency and parallelism. 
    * E.g. the concurrency primitives, the program’s runtime, and the operating system.
* Parallelism is a function of time/context.
    * For example, if our context was a space of five seconds, and we ran two operations that each took a second to run, we would consider the operations to have run in parallel. If our context was one second, we would consider the operations to have run sequentially.
* As we begin moving down the stack of abstraction, the problem of modeling things concurrently is becoming both more difficult to reason about, and more important.
* Most concurrent logic in our industry is written at one of the highest levels of abstraction: OS threads.
* In Go, we rarely have to think about our problem space in terms of OS threads. Instead, we model things in __goroutines__ and __channels__, and occasionally share memory.

## What Is CSP?

## How This Helps You?

## Go’s Philosophy on Concurrency

![Decision_tree_of_concurrency_primitives_and_channels](https://docs.google.com/drawings/d/e/2PACX-1vQvlFK3ViKVFBjrSD6BIT4A1cGxwdDJh_aO0sUeiDFgr0O-8EyEaVAdl3xqDVjoyEwRhfzxIHR4Q3KR/pub?w=825&h=557)