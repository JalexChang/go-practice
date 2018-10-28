# Notes

## Why Is Concurrency Hard?

Concurrent code is difficult to get right. It usually takes a few iterations to get it working as expected, and even then it’s not uncommon for bugs to exist in code for years before some change in timing causes a previously undiscovered bug to rear its head.

There are come commnon issues make working with concurrent code difficult:

### Race Conditions

* Race conditions are one of the most insidious types of concurrency bugs because they may not show up until years after the code has been placed into production.
* A race condition occurs when two or more operations must execute in the correct order, but the order is not promised by the program.
* It is also called a _data race_, where one concurrent operation attempts to read a variable while another concurrent operation is attempting to write to the same variable at the same time.
* Most of the time, data races are introduced because the developers are thinking about the problem sequentially. 

### Atomicity

* Having the property of atomicity means that within the context that it is operating, something is indivisible or uninterruptible.
* The atomicity of an operation can change depending on the currently defined scope; something may be atomic in one context, but not another. E.g. Operations are atomic within the context of a process may not be atomic in the context of the os.

### Memory Access Synchronization

* The scope in a program that needs exclusive access to a shared resource is called a _critical section_.
* The issue of memory access synchronization dicuss about how to manage critical sessions that help multiple processes can share resources to each other.
* It sometimes will create maintenance and performance problems.

### Deadlocks

* A deadlocked program is one in which all concurrent processes are waiting on one another. In this state, the program will never recover without outside intervention.
* A deadlock arises iff all of _Coffman Conditions_ hold simultaneously in a program:
    * Mutual Exclusion
    * Hold and Wait
    * No Preemption
    * Circular Wait

### Livelocks

* Livelocks are programs that are actively performing concurrent operations, but these operations do nothing to move the state of the program forward.
* A very common reason livelocks are written: two or more concurrent processes are attempting to prevent a deadlock without coordination. 

### Starvation

* Starvation is any situation where a concurrent process cannot get all the resources it needs to perform work.
* Starvation usually implies that there are one or more greedy concurrent process that are unfairly preventing one or more concurrent processes from accomplishing work as efficiently as possible.


### Determining Concurrency Safety

* The most difficulty of developing concurrent code is the thing that underlies all the other problems: people.

* If you are a developer and you are trying to wrangle all of these problems as you introduce new functionality, or fix bugs in your program, it can be really difficult to determine the right thing to do.
    * Who is responsible for the concurrency?
    * How is the problem space mapped onto concurrency primitives?
    * Who is responsible for the synchronization?


## Simplicity in the Face of Complexity

* The runtime and communication difficulties we’ve discussed are by no means solved by Go, but they have been made significantly easier.
* Go’s runtime does most of the heavy lifting and provides the foundation for most of Go’s concurrency niceties.
* Go’s runtime also automatically handles multiplexing concurrent operations onto operating system threads.
* Memory management can be another difficult problem domain in computer science, and when combined with concurrency, it can become extraordinarily difficult to write correct code.
* Go has made it much easier to use concurrency in your program by not forcing you to manage memory, let alone across concurrent processes.
    * As of Go 1.8, garbage collection pauses are generally between 10 and 100 ms.