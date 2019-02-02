# Notes

## Confinement

* Confinement is the simple but powerful idea of ensuring information is only available from one concurrent process.
    * There are two kinds of confinement possible: _ad hoc_ and _lexical_.
* _Ad hoc confinement_ is when you achieve confinement through a convention.
    * It is set by the languages community, the group you work within, or the codebase you work within.
    * Sticking to convention is difficult to achieve on projects of any size.
* _Lexical confinement_ involves using lexical scope to expose only the correct data and concurrency primitives for multiple concurrent processes to use.
    * Concurrent code that utilizes lexical confinement also has the benefit of usually being simpler to understand than concurrent code without lexically confined variables. 
* Synchronization comes with a cost, and if you can avoid it you won’t have any critical sections, and therefore you won’t have to pay the cost of synchronizing them.
* The confinement is difficult to establish.

## The for-select Loop

* The first variation keeps the _select_ statement as short as possible:
```
for {
    select {
    case <-done:
        return
    default:
    }

    // Do non-preemptable work
}
```

* The second variation embeds the work in a _default_ clause of the _select_ statement:
```
for {
    select {
    case <-done:
        return
    default:
        // Do non-preemptable work
    }
}
```

## Preventing Goroutine Leaks

* The runtime cannot garbage collect goroutines.
* A goroutine has a few paths to termination:
    * When it has completed its work.
    * When it cannot continue its work due to an unrecoverable error.
    * When it’s told to stop working.
* The parent goroutine can pass a channel to the child goroutine and then closes the channel when it wants to cancel the child goroutine.
* If a goroutine is responsible for creating a goroutine, it is also responsible for ensuring it can stop the goroutine.

## The or-channel

* Combine one or more done channels into a single done channel that closes if any of its component channels close.
* This pattern creates a composite __done__ channel through recursion and goroutines.
* Worrying about the number of goroutines created here is probably a premature optimization. 

## Error Handling

* Who should be responsible for handling the error?
    * Generally, it is __main goroutine__.
    * It has more context about the running program and can make more intelligent decision about what to do with errors.
* Don’t put your goroutines in a awkward position.
    * Only print the error and hope something is paying attention.
* Your concurrent processes should send their errors to another part of your program that has complete information about the state of your program.
    * This part of the program can make a more informed decision about what to do.
* Errors should be considered first-class citizens when constructing values to return from goroutines.
    *  If your goroutine can produce errors, those errors should be tightly coupled with your result type, and passed along through the same lines of communication.

## Pipelines

* A _pipeline_ is a powerful tool you can use to form an abstraction in your system for processing streams or batches of data.
    *  We call each of operations a __stage__ of the pipeline. 
* Benefits
    * Can modify stages independent of one another.
    * Can process each stage concurrent to upstream or downstream stages.
    * Can _fan-out_, or _rate-limit_ portions of your pipeline.
    * Can performs stream processing well.
* Properties of a pipeline stage
    * A stage consumes and returns the same type.
    * A stage must be reified so that it may be passed around.

### Best Practices for Constructing Pipelines

* Use Go’s _channel primitive
    * Each stage of the pipeline is executing concurrently.
    * How to promise the pipeline can be canceled?
        * Ranging over the incoming channel. When the incoming channel is closed, the range will exit.
        * The send sharing a select statement with the _done_ channel.

### Generators

* Converts a discrete set of values into a stream of data on a channel.
* __repeat__
    * Will repeat the values you pass to it infinitely until you tell it to stop.
* __take__
    * Will retrieve first N element in a channel.
* __repeat, take__
    * Can get N same elements.
* __repeatFn, take__
    * Can call a functions N times.
* It is OK to deal in channels of __interface{}__ so that you can use a standard library of pipeline patterns.
    * When you need to deal in specific types, you can place a stage that performs the __type assertion__ for you.
    * The type-specific stages are fast, but only marginally faster in magnitude.

## Fan-Out, Fan-In

* Sometimes, stages in your pipeline can be particularly computationally expensive.
    * Upstream stages in your pipeline can become blocked while waiting for your expensive stages to complete.
    * __Fan-out, Fan-in__ is a pattern that parallelize pulls from an upstream stage.
    * __Fan-out__ is a term to describe the process of starting multiple goroutines to handle input from the pipeline.
    * __Fan-in__ is a term to describe the process of combining multiple results into one channel.
* A naive implementation of the fan-in, fan-out algorithm only works if the order in which results arrive is unimportant.

## The or-done-channel

* A way to make your _select_ with _done_ much clear for readability.

## The tee-channel

* A way to send values coming in from a channel off into two separate areas of your codebase via different channels.
* The iteration over _in_ cannot continue until both _out1_ and _out2_ have been written to.

## The bridge-channel

* A way to consume values from a sequence of channels.
    * Ordered read from each channel.

## Queuing

* Queuing will almost _never speed up the total runtime_ of your program; it will only allow the program to behave differently.
    * Queuing is a way to reduce the time in _blocking state_.
    * The true utility of queues is to _decouple_ stages so that the runtime of one stage has no impact on the runtime of another.
* Only two situations that queuing can the overall performance of your system:
    * If _batching requests_ in a stage saves time.
        *  For example, a stage that buffers input in something faster (e.g., memory) than it is designed to send to (e.g., disk).
    * If _delays_ in a stage produce a feedback loop into the system.
        * This idea is often referred to as a _negative feedback loop_, _downward-spiral_, or even _death-spiral_. 
* Adding queuing prematurely can hide synchronization issues such as deadlocks and livelocks.
    * You may be tempted to add queuing elsewhere—e.g., after a computationally expensive stage—but avoid that temptation!
    * See _Little’s Law_.
* Queuing can be useful in your system, but because of its complexity, it’s usually one of the _last optimizations_.

## Context Package

* __Context__ pkg is a useful way to communicate extra information alongside the simple notification to cancel.
    * Why the cancellation was occuring.
    * Whether or not our function has a deadline by which it needs to complete.* Instances of __context.Context__ may look equivalent from the outside, but internally they may change at every stack-frame.
* Always pass instances of Context into your functions.
* Using the __done channel__ pattern, we could accomplish this by wrapping the incoming __done__ channel in other __done__ channels and then returning if any of them fire. But we wouldn’t have the extra information about __deadlines__ and __errors__ a __Context__ gives us.

### Package Methods

* __Context__ package provides serveral methods to help developers to control a goroutine's life cycle.
    * __Done__ method which returns a channel that’s closed when our function is to be preempted.
    * __Deadline__ function to indicate if a goroutine will be canceled after a certain time.
    * __Err__ method that will return non-nil if the goroutine was canceled.
    * __Value__ function provides request-specific information needs to be passed along in addition to information about preemption.
        * To provide a data-bag for transporting request-scoped data through your call-graph.
    * __WithCancel__ returns a new Context that closes its done channel when the returned __cancel__ function is called. 
    * __WithDeadline__ returns a new Context that closes its done channel when the machine’s clock advances past the given __deadline__.
    * __WithTimeout__ returns a new Context that closes its done channel after the given __timeout__ duration.
    * __WithValue__ returns a new Coontext that having a specific key-value in the data bag.
        * The __key__ you use must satisfy Go’s notion of _comparability_; that is, the equality operators _==_ and _!=_ need to return correct results when used.
        * __Values__ returned must be safe to access from multiple goroutines.


### Ruel of Thumb

* The Go authors recommend you follow a few rules when storing and retrieving value from a Context since the key-value is not type-safe.
    * Define a custom key-type in your package.
    * Use context values only for request-scoped data that transits processes and API boundaries, not for passing optional parameters to functions.
        * The data should transit process or API boundaries.
        * The data should be _immutable_.
        * The data should trend toward _simple types_.
        * The data should be data, _not types with methods_.
        * The data should _help decorate operations_, not drive them.
