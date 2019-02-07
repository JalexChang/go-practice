# Notes

## Work Stealing

* The algorithm used to handle multiplexing goroutines onto OS threads.
* In goroutines, _fair scheduling_ cannot execute well becuase Go models concurrency using a _fork-join_ model.
    * Tasks are likely dependent on one another. Will likely cause one of the processors to be underutilized.
    * Also lead to poor cache locality as tasks that require the same data are scheduled on other processors.
* _Centralized FIFO queue_ maybe can fix the underutilized problem, but it still have another problems.
    * Cannot support cache locality.
    * Cannot solve the case of fined-grained operations.
* __Decentralized work queue__
    * Give each processor its own thread and a _deque_.

### Stealing Tasks or Continuations?

* Under a fork-join paradigm, there are two options: __tasks__ and __continuations__.
    * Goroutines are tasks.
    * Everything after a goroutine is called is the continuation.
* Go’s work-stealing algorithm applies __continuation stealing__.
* The work stealing algorithm basic rules for __task stealing__:
    1. At a _fork_ point, add _tasks_ to the tail of the deque associated with the thread.
    2. If the thread is idle, steal work from the head of deque associated with some other random thread.
    3. At a _join_ point that cannot be realized yet, pop work off the tail of the thread’s own deque.
    4. If the thread’s deque is empty, either:
        1. Stall at a join.
        2. Steal work from the head of a random thread’s associated deque.
* The work stealing algorithm basic rules for __continuation stealing__:
    1. At a _fork_ point, add the _continuation_ to the tail of the deque associated with the thread and execute the _task_ immediately.
    2. If the thread is idle, steal work from the head of deque associated with some other random thread.
    3. At a _join_ point that cannot be realized yet, pop work off the tail of the thread’s own deque.
    4. If the thread’s deque is empty, either:
        1. Stall at a join.
        2. Steal work from the head of a random thread’s associated deque.
* Comparision:  
    || Continuation Stealing | Task Stealing (child) |
    |---|---|---|
    | Queue Size | Bounded | Unbounded |
    | Order of Execution | Serial | Out of Order |
    | Join Point | Nonstalling | Stalling |
* The drawback of continuation-stealing: requires support from the compiler.
* The work sitting on the tail of its deque has some interesting properties:
    * It’s the work most likely needed to complete the parent’s join (perform better).
    * It’s the work most likely to still be in our processor’s cache (fewer cache misses).
* For optimizing CPU utilization, Go supports a _global context_ (work queue) to steal back blocked goroutines.
    * Besides, when a context’s queue is empty, it will first check the global context for work to steal before checking other OS threads’ contexts.
* Related materials
    * [dotGo 2017 - JBD - Go's work stealing scheduler](https://www.youtube.com/watch?v=Yx6FBsGNOp4)
    * [GopherCon 2018: Kavya Joshi - The Scheduler Saga](https://www.youtube.com/watch?v=YHRO5WQGh0k)
