# Notes

## Goroutines

Every Go program has at least one goroutine: the _main_ goroutine, which is automatically created and started when the process begins.
* A goroutine is a function that is running concurrently.
* Can start a goroutine with the _go_ keyword.

Goroutines are not OS threads or exactly green threads (threads are managed by a language’s runtime), they are a higher level of abstraction known as _coroutines_.
* Coroutines are simply concurrent subroutines (functions, closures, or methods in Go).
* Coroutines are nonpreemptive. They cannot be interrupted but have multiple points which allow for suspension or reentry.
* Go’s runtime observes the runtime behavior of goroutines and automatically suspends them when they block and then resumes them when they become unblocked.

Go’s mechanism for hosting goroutines is an implementation of what’s called an _M:N scheduler_, which means it maps M green threads to N OS threads.

Go follows a model of concurrency called the _fork-join_ model.
* Fork: at any point in the program, it can split off a child branch of execution to be run concurrently with its parent.
* Join: at some point in the future, these concurrent branches of execution will join back together.

Goroutines operate within the same address space as each other, and simply host functions, utilizing goroutines is a natural extension to writing nonconcurrent code.
* Go’s compiler nicely takes care of pinning variables in memory so that goroutines don’t accidentally access freed memory
* Developers can focus on their problem space instead of memory management.
* But developers need to worry about synchronization when using goroutines.

Goroutines are lightweight that means they use less memory and have less context switch overhead cmparing to OS threads.
* Millions of goroutines without requiring swapping.
* 92% faster than an OS context switch (0.225 _us_)
