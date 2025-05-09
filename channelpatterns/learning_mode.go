package channelpatterns

import (
        "fmt"
        "strings"
)

// LearningModeStepExplanation contains detailed explanation for a step
type LearningModeStepExplanation struct {
        StepTitle       string
        Explanation     string
        CodeExample     string
        ExpectedOutput  string
        AdditionalNotes string
}

// GetLearningModeExplanation returns detailed step-by-step explanations for the specified pattern
func GetLearningModeExplanation(patternName string) []LearningModeStepExplanation {
        switch patternName {
        case "basic":
                return getBasicChannelExplanations()
        case "buffered":
                return getBufferedChannelExplanations()
        case "sync":
                return getSyncChannelExplanations()
        case "select":
                return getSelectExplanations()
        case "errors":
                return getErrorHandlingExplanations()
        case "pipeline":
                return getPipelineExplanations()
        default:
                return []LearningModeStepExplanation{
                        {
                                StepTitle:   "Unknown Pattern",
                                Explanation: "No detailed explanation available for " + patternName,
                        },
                }
        }
}

// DemonstrateLearningMode runs the specified pattern in learning mode with detailed explanations
func DemonstrateLearningMode(patternName string) {
        steps := GetLearningModeExplanation(patternName)
        
        fmt.Println("\n===== LEARNING MODE: " + strings.ToUpper(patternName) + " =====")
        fmt.Println("This mode provides step-by-step explanations of how this channel pattern works.")
        
        for i, step := range steps {
                fmt.Printf("\n--- STEP %d: %s ---\n\n", i+1, step.StepTitle)
                fmt.Println(step.Explanation)
                
                if step.CodeExample != "" {
                        fmt.Println("\nCode Example:")
                        fmt.Println(strings.Repeat("-", 40))
                        fmt.Println(step.CodeExample)
                        fmt.Println(strings.Repeat("-", 40))
                }
                
                if step.ExpectedOutput != "" {
                        fmt.Println("\nExpected Output:")
                        fmt.Println(strings.Repeat("-", 40))
                        fmt.Println(step.ExpectedOutput)
                        fmt.Println(strings.Repeat("-", 40))
                }
                
                if step.AdditionalNotes != "" {
                        fmt.Println("\nAdditional Notes:")
                        fmt.Println(step.AdditionalNotes)
                }
                
                // Pause between steps to allow reading
                if i < len(steps)-1 {
                        fmt.Println("\nPress Enter to continue to the next step...")
                        fmt.Scanln() // Wait for user to press Enter
                }
        }
        
        fmt.Println("\n===== LEARNING MODE COMPLETE =====")
        fmt.Println("You've completed the detailed explanation of " + patternName + " channel patterns.")
        fmt.Println("Next, you'll see the actual code in action.")
        fmt.Println("\nPress Enter to see the live demonstration...")
        fmt.Scanln() // Wait for user to press Enter
}

// GetBasicChannelExplanations returns explanations for basic channel operations
func getBasicChannelExplanations() []LearningModeStepExplanation {
        return []LearningModeStepExplanation{
                {
                        StepTitle: "Channel Basics",
                        Explanation: `Channels are a core feature of Go's concurrency model. They provide a way for 
goroutines to communicate with each other and synchronize their execution.

Channels are typed conduits through which you can send and receive values. The operator 
<- is used for sending and receiving data through channels.

To create a channel, you use the built-in make function:`,
                        CodeExample: `// Create an unbuffered channel that can send/receive string values
ch := make(chan string)

// Send a value into the channel (will block until someone receives)
ch <- "Hello"  

// Receive a value from the channel (will block until someone sends)
message := <-ch`,
                        AdditionalNotes: `Key characteristics of unbuffered channels:
1. Sending blocks until another goroutine receives
2. Receiving blocks until another goroutine sends
3. This creates a perfect synchronization point between goroutines`,
                },
                {
                        StepTitle: "Unidirectional Channels",
                        Explanation: `Channels can be restricted to only allow sending or only allow receiving.
This is useful when passing channels to functions - you can restrict what
operations the function can perform on the channel.

This provides additional type safety and makes your code's intent clearer.`,
                        CodeExample: `// A function that only sends to a channel (cannot receive)
func sender(ch chan<- int) {
    for i := 1; i <= 5; i++ {
        ch <- i
    }
    close(ch)
}

// A function that only receives from a channel (cannot send)
func receiver(ch <-chan int) {
    for num := range ch {
        fmt.Println("Received:", num)
    }
}`,
                        AdditionalNotes: `The channel directions are:
- chan<- T: Send-only channel (can only send)
- <-chan T: Receive-only channel (can only receive)
- chan T: Bidirectional channel (can send and receive)`,
                },
                {
                        StepTitle: "Channel Closing and Range",
                        Explanation: `Channels can be closed to signal that no more values will be sent.
This is important for letting receivers know when to stop waiting for values.

Receivers can check if a channel has been closed when receiving, or use
a for-range loop to automatically receive until the channel is closed.`,
                        CodeExample: `// Sender: Send values and then close
go func() {
    ch <- "Message 1"
    ch <- "Message 2"
    ch <- "Message 3"
    close(ch)  // Signal that no more messages will be sent
}()

// Method 1: Check if channel is closed while receiving
msg, ok := <-ch
if !ok {
    fmt.Println("Channel is closed")
}

// Method 2: Use range to receive until channel is closed
for msg := range ch {
    fmt.Println("Received:", msg)
}
// Loop exits automatically when channel is closed`,
                        ExpectedOutput: `Received: Message 1
Received: Message 2
Received: Message 3`,
                        AdditionalNotes: `Important rules about closing channels:
1. Only the sender should close a channel, never the receiver
2. Sending on a closed channel will cause a panic
3. Receiving from a closed channel returns the zero value immediately
4. The second return value from a receive tells you if the channel is still open`,
                },
        }
}

// GetBufferedChannelExplanations returns explanations for buffered channels
func getBufferedChannelExplanations() []LearningModeStepExplanation {
        return []LearningModeStepExplanation{
                {
                        StepTitle: "Buffered Channels",
                        Explanation: `Buffered channels have a capacity to hold multiple values before blocking.
Unlike unbuffered channels which block on send until a receiver is ready,
buffered channels only block when the buffer is full.

This allows for more flexible concurrency patterns with less synchronization.`,
                        CodeExample: `// Create a buffered channel with capacity of 2
bufferedCh := make(chan string, 2)

// Send values (won't block until buffer is full)
bufferedCh <- "First message"  // Doesn't block
bufferedCh <- "Second message" // Doesn't block
// bufferedCh <- "Third message"  // Would block until space is available

// Receive values from the buffer
fmt.Println(<-bufferedCh) // "First message"
fmt.Println(<-bufferedCh) // "Second message"`,
                        AdditionalNotes: `Buffered channels are useful when:
1. You want to reduce goroutine blocking
2. You know there might be bursts of messages
3. The producer and consumer run at different speeds
4. You want to limit the number of concurrent operations`,
                },
                {
                        StepTitle: "Buffer Capacity and Blocking",
                        Explanation: `A buffered channel has a fixed capacity defined when created. The behavior
depends on how full the buffer is:

1. When buffer is NOT FULL: Sends don't block
2. When buffer is FULL: Sends block until space becomes available
3. When buffer is NOT EMPTY: Receives don't block
4. When buffer is EMPTY: Receives block until values are sent

This creates backpressure when the buffer fills up, which can help
manage resource utilization.`,
                        CodeExample: `// Create a buffered channel with room for 2 items
ch := make(chan int, 2)

// Sending more items than the buffer capacity
go func() {
    for i := 1; i <= 5; i++ {
        fmt.Printf("Sending: %d\\n", i)
        ch <- i  // Will block after sending 2 items
        fmt.Printf("Sent: %d\\n", i)
    }
    close(ch)
}()

// Slow consumer that creates backpressure
// (In a real program, we would use time.Sleep here)
fmt.Println("Waiting 1 second to let sender fill buffer...")
// Slow receiver processing each item
for num := range ch {
    fmt.Printf("Received: %d\\n", num)
}`,
                        ExpectedOutput: `Sending: 1
Sent: 1
Sending: 2
Sent: 2
Sending: 3
[blocked...]
Received: 1
[unblocked]
Sent: 3
Sending: 4
[blocked...]`,
                        AdditionalNotes: `The capacity creates backpressure that propagates to upstream processes.
This is a key mechanism for managing load in concurrent systems.`,
                },
                {
                        StepTitle: "Choosing Between Buffered and Unbuffered Channels",
                        Explanation: `Deciding when to use buffered vs unbuffered channels depends on your concurrency needs:

UNBUFFERED CHANNELS:
- When you need strict synchronization between goroutines
- When you need to ensure each message is processed before sending the next
- When you want to ensure the sender knows the receiver has processed the item

BUFFERED CHANNELS:
- When senders and receivers work at different rates 
- When you want to batch process items
- When you want to limit concurrency (using the buffer as a semaphore)
- When you need to decouple producer and consumer timing`,
                        AdditionalNotes: `Performance note: While buffered channels can improve throughput in some scenarios, 
they don't necessarily make your program faster overall. They change the 
synchronization characteristics, which may or may not improve performance 
depending on your specific workload.`,
                },
        }
}

// GetSyncChannelExplanations returns explanations for synchronization patterns
func getSyncChannelExplanations() []LearningModeStepExplanation {
        return []LearningModeStepExplanation{
                {
                        StepTitle: "Channel Synchronization",
                        Explanation: `Channels can be used to synchronize execution between goroutines.
One of the simplest patterns is using a channel to signal when a goroutine 
has completed its work.

This is especially useful when you need to ensure a background task
has finished before continuing.`,
                        CodeExample: `// Create a channel to signal completion
done := make(chan bool)

// Start a goroutine and work
go func() {
    // Do some work...
    time.Sleep(2 * time.Second)
    
    // Signal that work is complete
    done <- true
}()

// Wait for the goroutine to finish
<-done  // This blocks until we receive a value
fmt.Println("Work complete, continuing execution")`,
                        AdditionalNotes: `This is a fundamental pattern in Go concurrency. You can use it to:
1. Wait for goroutines to finish
2. Signal between goroutines
3. Implement timeouts (combined with select)`,
                },
                {
                        StepTitle: "Worker Pools",
                        Explanation: `The worker pool pattern uses a fixed number of goroutines to process
work from a shared channel. This limits concurrency while maximizing throughput.

It consists of:
1. A channel for distributing work
2. A pool of worker goroutines
3. A channel for collecting results

This is one of the most common and useful channel patterns in Go.`,
                        CodeExample: `// Create channels for jobs and results
jobs := make(chan int, 100)
results := make(chan int, 100)

// Start 3 workers
for w := 1; w <= 3; w++ {
    go func(id int) {
        for job := range jobs {
            // Process job
            fmt.Printf("Worker %d processing job %d\\n", id, job)
            // Send result
            results <- job * 2
        }
    }(w)
}

// Send jobs
for j := 1; j <= 5; j++ {
    jobs <- j
}
close(jobs)  // No more jobs

// Collect results
for a := 1; a <= 5; a++ {
    result := <-results
    fmt.Println("Result:", result)
}`,
                        AdditionalNotes: `Worker pools provide several benefits:
1. Limit resource usage by controlling concurrency
2. Efficiently distribute work across multiple processors
3. Manage backpressure naturally through channel blocking
4. Allow processing items as soon as workers are available`,
                },
                {
                        StepTitle: "Fan-Out/Fan-In Pattern",
                        Explanation: `The fan-out/fan-in pattern distributes work across multiple goroutines
and then collects their results.

- Fan-out: Multiple goroutines read from a single channel, distributing work
- Fan-in: Results from multiple channels are combined into a single channel

This pattern is useful for CPU-intensive processing that can be parallelized.`,
                        CodeExample: `// Generator - produces values
func generator(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

// Fan-out - start multiple workers that process input
func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n  // Some CPU-intensive work
        }
        close(out)
    }()
    return out
}

// Fan-in - combine multiple result channels into one
func merge(cs ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)
    
    // Start an output goroutine for each input channel
    output := func(c <-chan int) {
        for n := range c {
            out <- n
        }
        wg.Done()
    }
    
    wg.Add(len(cs))
    for _, c := range cs {
        go output(c)
    }
    
    // Close when all output goroutines are done
    go func() {
        wg.Wait()
        close(out)
    }()
    
    return out
}

// Usage:
input := generator(1, 2, 3, 4, 5)

// Fan out to 3 workers
c1 := square(input)
c2 := square(input)
c3 := square(input)

// Fan in from 3 workers
for result := range merge(c1, c2, c3) {
    fmt.Println(result)
}`,
                        AdditionalNotes: `This pattern is ideal for:
1. CPU-bound tasks that benefit from parallel execution
2. I/O tasks where you want to make multiple requests simultaneously
3. When you need to process many items quickly by distributing the work`,
                },
                {
                        StepTitle: "Using Channels as Semaphores",
                        Explanation: `Buffered channels can act as semaphores to limit concurrent operations.
This is useful when you need to restrict access to limited resources.

The buffer size controls the maximum number of concurrent operations.`,
                        CodeExample: `// Create a semaphore channel with capacity 3
semaphore := make(chan struct{}, 3)

// Launch many goroutines that all want to run concurrently
var wg sync.WaitGroup
for i := 1; i <= 10; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        
        semaphore <- struct{}{}  // Acquire semaphore
        fmt.Printf("Task %d running\\n", id)
        time.Sleep(1 * time.Second)  // Simulate work
        <-semaphore  // Release semaphore
    }(i)
}

wg.Wait()  // Wait for all tasks to complete`,
                        AdditionalNotes: `This pattern is useful for:
1. Limiting concurrent database connections
2. Restricting network requests
3. Controlling access to limited resources
4. Implementing rate limiters

Note that we use 'struct{}{}' which takes 0 bytes of memory - perfect for
semaphores where we don't need to send actual data, just coordinate.`,
                },
        }
}

// GetSelectExplanations returns explanations for select pattern
func getSelectExplanations() []LearningModeStepExplanation {
        return []LearningModeStepExplanation{
                {
                        StepTitle: "Select Statement Basics",
                        Explanation: `The select statement lets you wait on multiple channel operations simultaneously.
It's similar to a switch statement but for channel operations.

Select blocks until one of its cases can proceed, then executes that case.
If multiple cases are ready, it chooses one at random.`,
                        CodeExample: `// Create two channels
ch1 := make(chan string)
ch2 := make(chan string)

// Send values on each channel asynchronously
go func() { ch1 <- "message from ch1" }()
go func() { ch2 <- "message from ch2" }()

// Wait for either channel to receive a value
select {
case msg1 := <-ch1:
    fmt.Println("Received from ch1:", msg1)
case msg2 := <-ch2:
    fmt.Println("Received from ch2:", msg2)
}`,
                        AdditionalNotes: `Select is one of the most powerful features in Go's concurrency model.
It allows you to:
1. Wait on multiple channels
2. Implement non-blocking operations
3. Implement timeouts
4. Combine multiple concurrent operations
5. Build more complex communication patterns`,
                },
                {
                        StepTitle: "Implementing Timeouts",
                        Explanation: `The select statement combined with time.After() provides an elegant way
to implement timeouts in concurrent operations.

This allows you to limit how long you'll wait for a channel operation,
which is essential for building responsive systems.`,
                        CodeExample: `// Create a channel for the operation
ch := make(chan string)

// Start an operation that might take too long
go func() {
    // Simulate work that takes 2 seconds
    time.Sleep(2 * time.Second)
    ch <- "operation complete"
}()

// Wait for the result with a timeout
select {
case result := <-ch:
    fmt.Println("Success:", result)
case <-time.After(1 * time.Second):
    fmt.Println("Timeout: operation took too long")
}`,
                        AdditionalNotes: `This pattern is essential for building resilient systems that don't
hang indefinitely. It's useful for:
1. Network requests
2. Database queries
3. User interactions
4. Any operation that shouldn't block forever

time.After() returns a channel that receives a value after the specified duration.`,
                },
                {
                        StepTitle: "Non-blocking Operations",
                        Explanation: `The select statement with a default case provides non-blocking channel operations.
This allows you to check if a channel is ready without blocking.

This is useful when you want to check a channel but continue execution
if the channel isn't ready yet.`,
                        CodeExample: `ch := make(chan string)

// Start a goroutine that will send a value after 1 second
go func() {
    time.Sleep(1 * time.Second)
    ch <- "delayed message"
}()

// Non-blocking receive attempt
select {
case msg := <-ch:
    fmt.Println("Received message:", msg)
default:
    fmt.Println("No message available, continuing execution")
}

// Later, we can try again...
time.Sleep(2 * time.Second)

select {
case msg := <-ch:
    fmt.Println("Received message:", msg)
default:
    fmt.Println("Still no message")
}`,
                        ExpectedOutput: `No message available, continuing execution
Received message: delayed message`,
                        AdditionalNotes: `This pattern is useful for polling channels or implementing high-performance
concurrent code that shouldn't block unnecessarily.

You can also use this pattern for non-blocking sends:

select {
case ch <- value:
    // Value was sent
default:
    // Channel wasn't ready for sending
}`,
                },
                {
                        StepTitle: "Channel Multiplexing",
                        Explanation: `Select can be used to multiplex messages from multiple channels onto a single
flow of execution. This allows you to process messages as they become available
from any of several channels.

This is useful when collecting results from parallel operations or handling
events from different sources.`,
                        CodeExample: `// Set up multiple channels
c1 := make(chan string)
c2 := make(chan string)
c3 := make(chan string)

// Send messages on each channel at different times
go func() {
    time.Sleep(100 * time.Millisecond)
    c1 <- "message from c1"
}()
go func() {
    time.Sleep(200 * time.Millisecond)
    c2 <- "message from c2"
}()
go func() {
    time.Sleep(300 * time.Millisecond)
    c3 <- "message from c3"
}()

// Receive messages in whatever order they arrive
for i := 0; i < 3; i++ {
    select {
    case msg1 := <-c1:
        fmt.Println(msg1)
    case msg2 := <-c2:
        fmt.Println(msg2)
    case msg3 := <-c3:
        fmt.Println(msg3)
    }
}`,
                        AdditionalNotes: `Multiplexing is particularly useful when:
1. Processing events from multiple sources
2. Combining results from parallel operations
3. Handling communication with multiple clients
4. Implementing event loops

For more advanced multiplexing, consider using the fan-in pattern
with a dedicated goroutine that combines multiple channels.`,
                },
                {
                        StepTitle: "Channel Closing Detection",
                        Explanation: `The select statement can detect when a channel is closed by using the
"comma ok" idiom in a case statement.

This allows you to handle channel closure as a specific event in your
select statement, rather than receiving the zero value.`,
                        CodeExample: `// Create a channel that will be closed
ch := make(chan int)

// Send values and then close
go func() {
    for i := 1; i <= 3; i++ {
        ch <- i
        time.Sleep(500 * time.Millisecond)
    }
    fmt.Println("Closing channel")
    close(ch)
}()

// Receive values and detect when channel is closed
for {
    select {
    case val, ok := <-ch:
        if !ok {
            fmt.Println("Channel is closed, exiting loop")
            return
        }
        fmt.Println("Received value:", val)
    }
}`,
                        ExpectedOutput: `Received value: 1
Received value: 2
Received value: 3
Closing channel
Channel is closed, exiting loop`,
                        AdditionalNotes: `The "comma ok" idiom (val, ok := <-ch) returns:
- val: The received value (zero value if channel is closed)
- ok: A boolean indicating if the channel is still open

This pattern is particularly important when you need to distinguish between:
1. Receiving a zero value from an open channel
2. Receiving the zero value because the channel is closed`,
                },
        }
}

// GetErrorHandlingExplanations returns explanations for error handling patterns
func getErrorHandlingExplanations() []LearningModeStepExplanation {
        return []LearningModeStepExplanation{
                {
                        StepTitle: "Error Handling with Result Channels",
                        Explanation: `In concurrent Go programs, you need to handle errors that occur in goroutines.
One common pattern is to use a "Result" struct with both a value and error field.

This allows you to pass both success values and errors through the same channel.`,
                        CodeExample: `// Define a Result type
type Result struct {
    Value int
    Err   error
}

// Function that returns a Result channel
func calculate(val int) <-chan Result {
    resultCh := make(chan Result)
    
    go func() {
        // Simulate work that might fail
        time.Sleep(time.Second)
        
        if val == 0 {
            // Return an error
            resultCh <- Result{Err: errors.New("cannot divide by zero")}
        } else {
            // Return a success value
            resultCh <- Result{Value: 100 / val}
        }
        close(resultCh)
    }()
    
    return resultCh
}

// Usage
for _, input := range []int{10, 0, 5} {
    result := <-calculate(input)
    if result.Err != nil {
        fmt.Printf("Error: %v\\n", result.Err)
    } else {
        fmt.Printf("Result: %d\\n", result.Value)
    }
}`,
                        ExpectedOutput: `Result: 10
Error: cannot divide by zero
Result: 20`,
                        AdditionalNotes: `This pattern is useful for any concurrent operation that can fail:
1. Network requests
2. Database operations 
3. File I/O
4. Complex calculations

Using a Result struct with both value and error fields follows Go's
standard error handling pattern in a concurrent context.`,
                },
                {
                        StepTitle: "Using Separate Error Channels",
                        Explanation: `Another approach to error handling in concurrent code is to use separate
channels for results and errors.

This can make sense when errors need special handling or when you want
to process either results or errors differently.`,
                        CodeExample: `// Function that returns separate value and error channels
func fetchData(query string) (<-chan string, <-chan error) {
    results := make(chan string)
    errors := make(chan error)
    
    go func() {
        // Simulate network request
        time.Sleep(time.Second)
        
        if query == "" {
            errors <- errors.New("empty query")
            close(results)
            close(errors)
            return
        }
        
        // Simulate successful response
        results <- "Data for: " + query
        close(results)
        close(errors)
    }()
    
    return results, errors
}

// Usage
results, errors := fetchData("example")

// Wait for either result or error
select {
case result := <-results:
    fmt.Println("Success:", result)
case err := <-errors:
    fmt.Println("Error:", err)
}`,
                        AdditionalNotes: `This pattern can be cleaner in certain situations:
1. When only a result OR an error is expected (not both)
2. When you want to handle errors on a separate path
3. When using fan-in/fan-out patterns with error aggregation

A variation is to use a context for cancellation combined with
result and error channels.`,
                },
                {
                        StepTitle: "Panic Recovery in Goroutines",
                        Explanation: `Panics in goroutines don't propagate to the parent goroutine - they will
crash the entire program if not recovered.

To handle panics in goroutines, you need to use defer and recover().`,
                        CodeExample: `// Create a channel to signal completion status
done := make(chan bool)

// Start a goroutine that might panic
go func() {
    // Set up panic recovery
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Recovered from panic: %v\\n", r)
            done <- false // Signal that we had an error
        }
    }()
    
    // This will cause a panic
    fmt.Println("About to panic...")
    panic("something went wrong")
    
    // This won't execute
    done <- true
}()

// Wait to see if the goroutine completed successfully
if <-done {
    fmt.Println("Goroutine completed successfully")
} else {
    fmt.Println("Goroutine had a problem")
}`,
                        ExpectedOutput: `About to panic...
Recovered from panic: something went wrong
Goroutine had a problem`,
                        AdditionalNotes: `It's critical to use panic recovery in all long-running goroutines or
goroutines that might panic to prevent your program from crashing.

Common panic scenarios in concurrent code:
1. Index out of range
2. Nil pointer dereference
3. Sending on a closed channel
4. Type assertion failures

After recovering, you can choose to:
1. Log the error
2. Return an error via a channel
3. Restart the goroutine
4. Continue with a fallback strategy`,
                },
                {
                        StepTitle: "Context-Based Cancellation",
                        Explanation: `The context package provides a standard way to propagate cancellation signals
to running goroutines.

This is especially important for:
1. Stopping long-running operations
2. Cleaning up resources
3. Preventing goroutine leaks`,
                        CodeExample: `// Function that supports cancellation via context
func worker(ctx context.Context) error {
    // Create a channel for reporting results
    resultCh := make(chan string)
    
    // Start the actual work in a goroutine
    go func() {
        // Simulate long task with periodic cancellation checks
        for i := 1; i <= 5; i++ {
            // Check for cancellation before each unit of work
            select {
            case <-ctx.Done():
                return // Exit early if cancelled
            default:
                // Continue working
            }
            
            // Do some work
            time.Sleep(500 * time.Millisecond)
            fmt.Printf("Step %d complete\\n", i)
        }
        
        resultCh <- "operation successful"
    }()
    
    // Wait for either completion or cancellation
    select {
    case result := <-resultCh:
        return nil
    case <-ctx.Done():
        return ctx.Err() // Return the cancellation reason
    }
}

// Usage
ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
defer cancel()

if err := worker(ctx); err != nil {
    fmt.Printf("Operation failed: %v\\n", err)
} else {
    fmt.Println("Operation succeeded")
}`,
                        ExpectedOutput: `Step 1 complete
Step 2 complete
Step 3 complete
Operation failed: context deadline exceeded`,
                        AdditionalNotes: `Context-based cancellation is the standard way to cancel operations in Go.
It allows cancellation to propagate through call stacks and across API boundaries.

Common ways to create cancellable contexts:
1. context.WithCancel(parent): Manual cancellation
2. context.WithTimeout(parent, duration): Time-based cancellation
3. context.WithDeadline(parent, time): Deadline-based cancellation

Always remember to call the cancel function when you're done with the context,
even if it has already expired, to free up resources.`,
                },
                {
                        StepTitle: "Deadlock Avoidance",
                        Explanation: `Deadlocks occur when goroutines are permanently blocked waiting for each other.
With channels, this typically happens when:

1. All goroutines are blocked on channel operations
2. No goroutine can proceed because they're all waiting

Here are key strategies to avoid deadlocks:`,
                        CodeExample: `// Strategy 1: Use buffered channels
ch := make(chan int, 1) // Buffer of 1
ch <- 1  // Won't block
x := <-ch  // Now we can receive

// Strategy 2: Use select with default to avoid blocking
select {
case ch <- value:
    // Value sent
default:
    // Channel wasn't ready, do something else
}

// Strategy 3: Always ensure proper channel closing
go func() {
    for i := 1; i <= 5; i++ {
        ch <- i
    }
    close(ch)  // Important!
}()

// Strategy 4: Use timeouts to prevent permanent blocking
select {
case res := <-ch:
    // Handle result
case <-time.After(1 * time.Second):
    // Timed out, handle accordingly
}`,
                        AdditionalNotes: `Common deadlock causes:
1. Forgetting to close channels
2. Incorrect channel direction (receive-only vs send-only)
3. Circular dependencies between goroutines
4. Incorrect use of mutexes together with channels

Go's runtime will detect some deadlocks automatically, but
only in situations where all goroutines are blocked.`,
                },
        }
}

// GetPipelineExplanations returns explanations for the data pipeline pattern
func getPipelineExplanations() []LearningModeStepExplanation {
        return []LearningModeStepExplanation{
                {
                        StepTitle: "Data Processing Pipelines",
                        Explanation: `A data pipeline processes data through a series of stages, where each stage
receives data from the previous stage, performs some processing, and sends
the result to the next stage.

Channels are ideal for building pipelines because they naturally represent 
this flow of data between processing stages.`,
                        CodeExample: `// A simple three-stage pipeline
//
//    Generate Numbers → Square Numbers → Sum Results
//
func generateNumbers(max int) <-chan int {
    out := make(chan int)
    go func() {
        for i := 1; i <= max; i++ {
            out <- i
        }
        close(out)
    }()
    return out
}

func squareNumbers(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

func sumResults(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        sum := 0
        for n := range in {
            sum += n
        }
        out <- sum
        close(out)
    }()
    return out
}

// Usage
numbers := generateNumbers(5)       // Stage 1: Generate numbers 1-5
squares := squareNumbers(numbers)   // Stage 2: Square each number
sum := sumResults(squares)          // Stage 3: Sum the squares

fmt.Println("Sum of squares:", <-sum)`,
                        ExpectedOutput: `Sum of squares: 55  // (1² + 2² + 3² + 4² + 5² = 1 + 4 + 9 + 16 + 25 = 55)`,
                        AdditionalNotes: `Key benefits of channel-based pipelines:
1. Each stage runs concurrently in its own goroutine
2. Stages are decoupled and focused on a single task
3. Data flows naturally through the system
4. Backpressure is handled automatically through channel blocking
5. Pipeline stages can be composed and reused`,
                },
                {
                        StepTitle: "Fan-Out in Pipelines",
                        Explanation: `For CPU-intensive stages in a pipeline, we can "fan out" the work to
multiple goroutines to process items in parallel.

This is particularly useful for stages that:
1. Are computationally expensive
2. Can process items independently
3. Would benefit from parallel execution`,
                        CodeExample: `// A pipeline with fan-out in the processing stage:
//
//               ┌→ Process →┐
// Generate →───┼→ Process →┼→ Collect
//               └→ Process →┘
//
func generateWork(count int) <-chan int {
    out := make(chan int)
    go func() {
        for i := 0; i < count; i++ {
            out <- i
        }
        close(out)
    }()
    return out
}

func processWork(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            // Simulate expensive computation
            time.Sleep(100 * time.Millisecond)
            out <- n * n
        }
        close(out)
    }()
    return out
}

func fanOut(in <-chan int, numWorkers int) []<-chan int {
    outputs := make([]<-chan int, numWorkers)
    for i := 0; i < numWorkers; i++ {
        outputs[i] = processWork(in) // Each worker shares the input
    }
    return outputs
}

func merge(cs ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)
    
    for _, c := range cs {
        wg.Add(1)
        go func(ch <-chan int) {
            defer wg.Done()
            for n := range ch {
                out <- n
            }
        }(c)
    }
    
    go func() {
        wg.Wait()
        close(out)
    }()
    
    return out
}

// Usage
const numJobs = 10
const numWorkers = 3

input := generateWork(numJobs)
workers := fanOut(input, numWorkers)
results := merge(workers...)

// Collect all results
for result := range results {
    fmt.Println(result)
}`,
                        AdditionalNotes: `The fan-out pattern is ideal for CPU-bound operations where
parallelism can improve throughput.

Things to consider:
1. The optimal number of workers depends on your CPU count
2. Too many workers can lead to excessive context switching
3. The work should be independent (no shared state)`,
                },
                {
                        StepTitle: "Batch Processing in Pipelines",
                        Explanation: `Sometimes it's more efficient to process items in batches rather than
one at a time. This is common for:

1. Database operations (bulk inserts)
2. API calls (batch requests)
3. I/O operations (bulk file writes)

Channels can be used to implement batch processing in pipelines.`,
                        CodeExample: `// A pipeline with batch processing:
//
// Generate → Batch → Process Batches → Results
//
func generateItems(count int) <-chan string {
    out := make(chan string)
    go func() {
        for i := 1; i <= count; i++ {
            out <- fmt.Sprintf("Item-%d", i)
            time.Sleep(100 * time.Millisecond)
        }
        close(out)
    }()
    return out
}

func batch(in <-chan string, size int) <-chan []string {
    out := make(chan []string)
    go func() {
        batch := make([]string, 0, size)
        
        for item := range in {
            batch = append(batch, item)
            
            // When batch is full or input is done, send the batch
            if len(batch) >= size {
                out <- batch
                batch = make([]string, 0, size)
            }
        }
        
        // Send any remaining items in the last batch
        if len(batch) > 0 {
            out <- batch
        }
        
        close(out)
    }()
    return out
}

func processBatch(in <-chan []string) <-chan string {
    out := make(chan string)
    go func() {
        for itemBatch := range in {
            // Process the entire batch at once
            fmt.Printf("Processing batch of %d items\\n", len(itemBatch))
            time.Sleep(500 * time.Millisecond)
            
            // Output individual results
            for _, item := range itemBatch {
                out <- fmt.Sprintf("Processed-%s", item)
            }
        }
        close(out)
    }()
    return out
}

// Usage
items := generateItems(10)
batches := batch(items, 3)
results := processBatch(batches)

// Collect and print results
for result := range results {
    fmt.Println(result)
}`,
                        AdditionalNotes: `Batch processing is particularly useful when:
1. There's overhead for each operation (network, disk)
2. Throughput is more important than latency
3. The batch operation is more efficient than individual operations

You can combine batching with other patterns like fan-out to 
process multiple batches in parallel.`,
                },
                {
                        StepTitle: "Error Handling in Pipelines",
                        Explanation: `Real-world pipelines need robust error handling. There are several approaches:

1. Result struct with Value and Error fields
2. Separate error channels
3. Error aggregation from multiple stages

The design depends on how you want to handle errors:
- Fail fast on first error
- Collect and report all errors
- Partial success with some errors`,
                        CodeExample: `// A pipeline with error handling:
//
// Generate → Process (with errors) → Filter Errors → Valid Results
//

// Result type that can contain either data or an error
type Result struct {
    Data  string
    Error error
}

func generateData(count int) <-chan string {
    out := make(chan string)
    go func() {
        for i := 1; i <= count; i++ {
            out <- fmt.Sprintf("data-%d", i)
        }
        close(out)
    }()
    return out
}

func processWithErrors(in <-chan string) <-chan Result {
    out := make(chan Result)
    go func() {
        for data := range in {
            // Simulate occasional errors
            if rand.Intn(3) == 0 {
                out <- Result{Error: fmt.Errorf("error processing %s", data)}
            } else {
                out <- Result{Data: "processed-" + data}
            }
        }
        close(out)
    }()
    return out
}

func filterErrors(in <-chan Result) (<-chan string, <-chan error) {
    outData := make(chan string)
    outErr := make(chan error)
    
    go func() {
        for result := range in {
            if result.Error != nil {
                outErr <- result.Error
            } else {
                outData <- result.Data
            }
        }
        close(outData)
        close(outErr)
    }()
    
    return outData, outErr
}

// Usage
input := generateData(10)
processed := processWithErrors(input)
validResults, errors := filterErrors(processed)

// Handle errors in a separate goroutine
go func() {
    for err := range errors {
        fmt.Printf("ERROR: %v\\n", err)
    }
}()

// Process valid results
for data := range validResults {
    fmt.Printf("SUCCESS: %s\\n", data)
}`,
                        AdditionalNotes: `When designing error handling for pipelines, consider:
1. Should errors stop the entire pipeline or just skip items?
2. Do you need to collect all errors or just report them?
3. Are some errors recoverable?
4. Do you need to associate errors with specific items?

For critical systems, you might also need retry logic and
circuit breakers to handle transient failures.`,
                },
        }
}