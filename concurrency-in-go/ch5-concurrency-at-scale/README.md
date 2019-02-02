# Notes

## Error Propagation

* Many developers make the error propagation as secondary, or “other,” to the flow of their system. 
* Few pieces of critical information an error should tell us:
    * __What happened__
    * __When and where it occurred__
        * Errors should always contain a _complete stack trace_.
    * A friendly __user-facing message__
        * The message that gets displayed to the user should be customized to suit your system and its users.
    * How the user can __get more information__.
        * E.g. Provide a ID to cross-reference to a corresponding log that displays the full information of the error, such as time the error occurred and the stack trace.
*  Place all errors into one of two categories: _bugs_ and _known cases_.
    * __Bugs__ are errors that you have not customized to your system.
    * __Known cases__ are edge cases you have algready known and have handle them by customized ways.
        * E.g. broken network connections, failed disk writes, etc.
* When our user-facing code receives a _well-formed error_:
    * We can be confident that care was taken at all levels in our code.
    * We can simply log it and print it out for the user to see.
    * Example of a well-formed error:
    ```go
    func PostReport(id string) error {
        result, err := lowlevel.DoWork()
        if err != nil {
            if _, ok := err.(lowlevel.Error); ok {
                err = WrapErr(err, "cannot post report with id %q", id)
            }
            return err
        }
    // ...
    }
    ```

## Timeouts and Cancellation

### When would concurrent processes need timeouts?
    
* System saturation
    * If our system is saturated, we may want requests at the edges of our system to time out rather than take a long time to field them.
    * Use case
        * The request is unlikely to be repeated when it is timed out.
        * You don’t have the resources to store the requests.
        * If the need for the request, or the data it’s sending, will go stale.
* Stale data
    * Data has a window within which it must be processed before more relevant data is available, or the need to process the data has expired.
* Deadlock
    * It is recommended to place timeouts on all of your concurrent operations to guarantee your system won’t deadlock.
    * By setting a timeout, the system can potentially transform your problem from _deadlocks_ to _livelocks_ which are much easier to solve.

### Why a concurrent process might be canceled?

* Timeouts
* User intervention
    * For user-facing concurrent operations, it is sometimes necessary to allow the users to cancel the operation they’ve started.
* Parent cancellation
    * If any kind of parent of a concurrent operation stops, the child will be canceled.
* Replicated requests
    * We may send data to multiple concurrent processes in an attempt to get a faster response from one of them. When the first one comes back, we would want to cancel the rest of the processes.

### What things do you need to take into account when writing concurrent code that can be terminated at any time?

* Define the period within which our concurrent process is preemptable.
* Ensure that any functionality that takes more time than this period is itself _preemptable_.
* An easy way to do this is to _break up_ the pieces of your goroutine into _smaller pieces_.

### If our goroutine happens to modify shared state (e.g. databasesa and data structures), what happens when the goroutine is canceled?

* Keep your modifications to any shared state within a tight scope.
* Ensure those modifications are easily rolled back.

### How can I solve the problem of duplicate messages caused by cancellation in pipelines?

* __Heartbeat__ (recommended)
    * A parent goroutine will send a cancellation signal after a child goroutine has already reported a result.
* Accept either the first or last result reported.
    * If your concurrent process is _idempotent_.
*  Poll the parent goroutine for permission
    * Use bidirectional communication with your parent to explicitly request permission to send your message.
    * It is similar to but more complicated than heartbeats.

## Heartbeat

* __Heartbeat__ is a way for concurrent processes to signal life to outside parties.
    * There are two types: __occurring on a time interval__ and __occurring at the beginning of a unit of work__.

### Heartbeat occurring on a time interval

* It is a way to periodically signal to its listeners that everything is well.   
* It is useful for concurrent code that might be waiting for something else to happen for it to process a unit of work.

### Heartbeat occurring at the beginning of a unit of work

* If you only care that the goroutine has started doing its work, this style of heartbeat is simple.
* This style can also help us _test multi-goroutine case_.
    * Instead of using _timeout_ to indicate whether the goroutine has finish its job or not, we can wait a signal from heartbeat channel, which is _deterministic_.
    * By the way, we can safely write our test without timeouts.
* If you’re reasonably sure the goroutine’s loop won’t stop executing once it is started, _only blocking on the first heartbeat_ is recommended.

## Replicated Requests

* We may send multiple replicated requests in an attempt to get a faster response from one of them.
    * The downside is that you’ll have to utilize resources to keep multiple copies of the handlers running.
* Although this is can be expensive to set up and maintain, if speed is your goal, this is a valuable technique.
* In addition, this naturally provides fault tolerance and scalability.

## Rate Limiting

* Rate limiting allows you to reason about the performance and stability of your system by preventing it from falling outside the boundaries.
* By rate limiting a system, you prevent entire classes of attack vectors against your system.
* Rate limits are often thought of from the perspective of people who _build_ the resources being limited, but rate limiting can also be utilized by users.
    * e.g. servers' API
* Most rate limiting is done by utilizing an algorithm called the __token bucket__.
    * [Leaky Bucket vs Tocken Bucket](https://www.slideshare.net/vimal25792/leaky-bucket-tocken-buckettraffic-shaping)
* Useful pkg: [golang.org/x/time/rate ](golang.org/x/time/rate)

### Multi-limiter
 
* We will probably want to establish multiple tiers of limits: coarse-grained controls to limit requests per second, minute, hour, and day.
* The easier way is to keep the limiters separate and then combine them into one rate limiter. Called __multiLimiter__.
* We will _sort_ the child RateLimiter instances when multiLimiter is instantiated.
    * It can simply return the most restrictive limit, which will be the first element in the slice.

## Healing Unhealthy Goroutines

* It can be very easy for a goroutine to become stuck in a bad state from which it cannot recover _without external help._
* In a long-running process, it can be useful to create a mechanism that ensures your goroutines remain healthy and restarts them if they become unhealthy.
    * _Heartbeat_ pattern is used to check up on the liveliness of the goroutine we’re monitoring.
* _Steward_: the logic that monitors a goroutine’s health.
    * It needs a reference to a function that can start the _ward_.
* _Ward_: the goroutine that a _steward_ monitors.
