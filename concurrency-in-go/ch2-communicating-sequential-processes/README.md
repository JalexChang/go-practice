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

* CSP stands for “Communicating Sequential Processes,” which is both a technique and the name of the paper (_Charles Antony Richard Hoare_, 1978) that introduced it.
* Hoare suggests that input and output are two overlooked primitives of programming—particularly in concurrent code.
* For communication between the processes, Hoare created input and output commands: __!__ for sending input into a process, and __?__ for reading output from a process. The operations and rules are apparently similar to to Go’s channels.
* Memory access synchronization isn’t inherently bad; however, the shared memory model can be difficult to utilize correctly—especially in large or complicated programs.
* Go has been built from the start with principles from CSP in mind and therefore it is easy to read, write, and reason about.

## How This Helps You?

* Goroutines free us from having to think about our problem space in terms of parallelism and instead allow us to model problems closer to their natural level of concurrency.
* For example, there are some common issues about concurrency to arise when building a web server:
    * Does my language naturally support threads, or will I have to pick a library?
    * Where should my thread confinement boundaries be?
    * How heavy are threads in this operating system?
    * How do the operating systems my program will be running in handle threads differently?
    * I should create a pool of workers to constrain the number of threads I create. How do I find the optimal number?
* Goroutines are lightweight, and we normally won’t have to worry about creating one.
    * There are appropriate times to consider how many goroutines are running in your system, but doing so upfront is soundly a premature optimization. 
* Go decouples concurrency and parallelism. 
    * Go’s runtime is managing the scheduling of goroutines for you, it can introspect on things like goroutines blocked waiting for I/O and intelligently reallocate OS threads to goroutines that are not blocked.
* Another benefit is that Go let us work on problems in naturally concurrent logic, we’ll naturally be writing concurrent code at a finer level of granularity than we perhaps would in other languages.

## Go’s Philosophy on Concurrency

![Decision_tree_of_concurrency_primitives_and_channels](https://docs.google.com/drawings/d/e/2PACX-1vQvlFK3ViKVFBjrSD6BIT4A1cGxwdDJh_aO0sUeiDFgr0O-8EyEaVAdl3xqDVjoyEwRhfzxIHR4Q3KR/pub?w=825&h=557)

* Is it a performance-critical section?
    * Channels use memory access synchronization to operate, therefore they can only be slower.
    * A performance-critical section might be hinting that we need to restructure our program.
* Are you trying to transfer ownership of data?
    * Data has an owner, and one way to make concurrent programs safe is to ensure only one concurrent context has ownership of data at a time.
    * Also, we can create buffered channels to implement a cheap in-memory queue.
* Are you trying to guard internal state of a struct?
    * By using memory access synchronization primitives, you can hide the implementation detail of locking your critical section from your callers.
* Are you trying to coordinate multiple pieces of logic?
    * Channels are inherently more composable than memory access synchronization primitives. 
* Conclusion
    * Aim for simplicity, use channels when possible, and treat goroutines like a free resource.