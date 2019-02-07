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
* The work stealing algorithm follows a few basic rules:
    1. At a fork point, add tasks to the tail of the deque associated with the thread.
    2. If the thread is idle, steal work from the head of deque associated with some other random thread.
    3. At a join point that cannot be realized yet, pop work off the tail of the thread’s own deque.
    4. If the thread’s deque is empty, either:
        1. Stall at a join.
        2. Steal work from the head of a random thread’s associated deque.
