# Channels

## What are 'channels' in Golang

Channels in Golang are the primary way for goroutines to communicate and synchronize with each other. They are typed conduits through which you can send and receive values with the channel operator, `<-`. This communication is blocking by default, which means a sender will wait for a receiver and vice-versa, allowing for safe and controlled data exchange between concurrent processes without explicit locks.

```go
package main

import "fmt"

func main() {
    // Create a new channel of type string.
    messages := make(chan string)

    // Start a new goroutine that sends a message to the channel.
    go func() { messages <- "ping" }()

    // Receive the message from the channel.
    msg := <-messages
    fmt.Println(msg)
}
```

## What is 'concurrency' in Golang

Concurrency in Golang is the ability to have multiple tasks running independently, but not necessarily simultaneously. It's achieved through the use of "goroutines," which are lightweight, independently executing functions. These goroutines can communicate and synchronize with each other using "channels," which are typed conduits for sending and receiving data. This model of "communicating sequential processes" allows for building complex concurrent applications in a safe and structured way, avoiding many of the pitfalls of traditional multi-threaded programming with shared memory.

```go
package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    go say("world")
    say("hello")
}
```

## What is a 'deadlock'

A deadlock in Golang is a situation where two or more goroutines are blocked forever, waiting for each other to release a resource that they need. This typically happens when goroutines are trying to acquire locks on shared resources in a different order, or when they are trying to send or receive data from channels in a way that creates a circular dependency. The Go runtime can detect when all goroutines are asleep and will panic with a "fatal error: all goroutines are asleep - deadlock!" message.

```go
package main

import "fmt"

func main() {
    // Create a channel.
    ch := make(chan int)

    // This will cause a deadlock because there is no other goroutine
    // to receive the data from the channel.
    ch <- 1

    fmt.Println("This will never be printed.")
}
```

## Building an understanding of processing approaches

Feature | Concurrency | Parallelism | Actor Model |
|---|---|---|---|
| **Definition** | Manages multiple tasks at once, but not necessarily executing them simultaneously. | Executes multiple tasks simultaneously. | A model of concurrent computation where independent "actors" communicate through asynchronous messages. |
| **Execution** | Tasks can be interleaved, giving the illusion of simultaneous execution. | Tasks run at the same time on multiple cores or processors. | Each actor processes messages sequentially, but multiple actors can run concurrently or in parallel. |
| **Resource Sharing** | Often involves shared memory and requires synchronization mechanisms like locks to prevent race conditions. | Can use shared memory or distributed memory, depending on the architecture. | No shared memory between actors. State is encapsulated within each actor. |
| **Communication** | Can be through shared memory, message passing, or other inter-process communication mechanisms. | Typically through shared memory or message passing interfaces (like MPI). | Exclusively through asynchronous message passing. |
| **Key Benefit** | Improved responsiveness and resource utilization, especially for I/O-bound tasks. | Significant performance gains for CPU-bound tasks that can be broken down into independent sub-tasks. | High level of concurrency with no risk of race conditions, excellent for building scalable and fault-tolerant systems. |
| **Risks/Challenges**| Deadlocks, race conditions, and other synchronization issues if not managed carefully. | Can be complex to program and debug. Not all problems can be parallelized effectively. | Can be more complex to reason about due to its asynchronous nature. Message passing overhead can be a factor.
