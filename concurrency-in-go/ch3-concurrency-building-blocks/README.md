# Notes

## Goroutines

* Every Go program has at least one goroutine: the _main_ goroutine, which is automatically created and started when the process begins.
    * A goroutine is a function that is running concurrently.
    * Can start a goroutine with the _go_ keyword.

* Goroutines are not OS threads or exactly green threads (threads are managed by a language’s runtime), they are a higher level of abstraction known as _coroutines_.
    * Coroutines are simply concurrent subroutines (functions, closures, or methods in Go).
    * Coroutines are nonpreemptive. They cannot be interrupted but have multiple points which allow for suspension or reentry.
    * Go’s runtime observes the runtime behavior of goroutines and automatically suspends them when they block and then resumes them when they become unblocked.

* Go’s mechanism for hosting goroutines is an implementation of what’s called an _M:N scheduler_, which means it maps M green threads to N OS threads.

* Go follows a model of concurrency called the _fork-join_ model.
    * Fork: at any point in the program, it can split off a child branch of execution to be run concurrently with its parent.
    * Join: at some point in the future, these concurrent branches of execution will join back together.

* Goroutines operate within the same address space as each other, and simply host functions, utilizing goroutines is a natural extension to writing nonconcurrent code.
    * Go’s compiler nicely takes care of pinning variables in memory so that goroutines don’t accidentally access freed memory
    * Developers can focus on their problem space instead of memory management.
    * But developers need to worry about synchronization when using goroutines.

* Goroutines are lightweight that means they use less memory and have less context switch overhead cmparing to OS threads.
    * Millions of goroutines without requiring swapping.
    * 92% faster than an OS context switch (0.225 _us_)

## WaitGroup

* WaitGroup is a great way to wait for a set of concurrent operations to complete when you either don’t care about the result of the concurrent operation, or you have other means of collecting their results.
    * Calls to __Add__ increment the counter by the integer passed in.
    * Calls to __Done__ decrement the counter by one.
    * Calls to __Wait__ block until the counter is zero.

* The calls to __Add__ are done outside the goroutines they’re helping to track. If we didn’t do this, we would have introduced a race condition.

## Mutex and RWMutex

* __Mutex__ provides a concurrent-safe way to express exclusive access to these shared resources.
    *  Should always call __Unlock__ within a _defer_ statement.
* __RWMutex__ gives you a little bit more control over the memory.
    * You can request a lock for reading, in which case you will be granted access unless the lock is being held for writing.

## Cond

* Ｕse __Cond__ if you want to wait for a signal before continuing execution on a goroutine.
* The call to __Wait__ doesn’t just block, it suspends the current goroutine, allowing other goroutines to run on the OS thread. 
    * When you call __Wait__: upon entering Wait, Unlock is called on the Cond variable’s Locker, and upon exiting Wait, Lock is called on the Cond variable’s Locker. 
* Internally, the runtime maintains a FIFO list of goroutines waiting to be signaled; __Signal__ finds the goroutine that’s been waiting the longest and notifies that, whereas __Broadcast__ sends a signal to all goroutines that are waiting.

## Once

* sync.Once is a type that utilizes some sync primitives internally to ensure that only one call to Do ever calls the function passed in—even on different goroutines.

## Pool

* Pool is a concurrent-safe implementation of the object pool pattern.
* Pool’s primary interface is its Get method. 
    * When called, Get will first check whether there are any available instances within the pool to return to the caller
    * If not, call its New member variable to create a new one. 
    When finished, callers call Put to place the instance they were working with back in the pool for use by other processes.
* Another common situation where a Pool is useful is for warming a cache of pre-allocated objects for operations that must run as quickly as possible. 
    * This is very common when writing high-throughput network servers that attempt to respond to requests as quickly as possible.
* Thumb rules
    * When instantiating sync.Pool, give it a New member variable that is thread-safe when called.
    * When you receive an instance from Get, make no assumptions regarding the state of the object you receive back.
    * Make sure to call Put when you’re finished with the object you pulled out of the pool. Otherwise, the Pool is useless. Usually this is done with defer.
    * Objects in the pool must be roughly uniform in makeup.

## Channels

* Channels are one of the synchronization primitives in Go derived from Hoare’s CSP. 
    * They are best used to communicate information between goroutines.
* Buffered channels can be useful in certain situations, but you should create them with care.
    * Buffered channels can easily become a premature optimization and also hide deadlocks by making them more unlikely to happen.
    * Take a good look at the example [314](./channels/314-using-buffered-chans.go).
* Result of channel operations given a channel’s state

    | Operation | Channel state | Result | 
    |---|---|---|
    | Read | nil | Block | 
    | | Open and Not Empty | Value |
    | | Open and Empty | Block |
    | | Closed | \<default value>, false |
    | | Write Only | Compilation Error |
    | Write | nil | Block | 
    | | Open and Full | Block |
    | | Open and Not Full | Write Value |
    | | Closed | __panic__ |
    | | Receive Only | Compilation Error |
    | Close | nil | __panic__ | 
    | | Open and Not Empty | Closes Channel; reads succeed until channel is drained, then reads produce default value |
    | | Open and Empty | Closes Channel; reads produces default value |
    | | Closed | __panic__ |
    | | Receive Only | Compilation Error |

* The first thing we should do to put channels in the right context is to assign channel _ownership_.
    * Define ownership as being a goroutine that instantiates, writes, and closes a channel.

