# Go Channel Patterns

This project demonstrates various channel patterns and techniques for effective concurrency in Go. It provides practical examples showing how channels work in different scenarios.

## What are Go Channels?

Channels are a core feature of Go's concurrency model. They provide a way for goroutines to communicate with each other and synchronize their execution without explicit locks or condition variables.

## Running the Examples

To run all examples:

```bash
go run main.go
```

The program will automatically cycle through all examples with a brief pause between each demonstration.

## Learning Mode with Step-by-Step Explanations

The program includes a special learning mode that provides detailed step-by-step explanations of each channel pattern before demonstrating it.

To see available learning topics:

```bash
go run main.go -list
```

To run the learning mode for a specific topic:

```bash
go run main.go -learn basic
go run main.go -learn buffered
go run main.go -learn sync
go run main.go -learn select
go run main.go -learn errors
go run main.go -learn pipeline
```

Learning mode features:
- Detailed explanations of concepts
- Code examples with expected output
- Additional notes and best practices
- Live demonstration after the explanations

## Examples Included

1. **Basic Channel Communication**
   - Simple send/receive operations
   - Unidirectional channels
   - Channel iteration and closing

2. **Buffered vs Unbuffered Channels**
   - Differences in behavior
   - Buffer capacity and blocking

3. **Channel Synchronization**
   - Simple synchronization
   - Worker pool pattern
   - Fan-out/fan-in pattern
   - Using channels as semaphores

4. **Select Statement**
   - Waiting on multiple channels
   - Timeouts
   - Non-blocking operations with default case
   - Multiplexing multiple channels
   - Detecting when a channel is closed

5. **Error Handling with Channels**
   - Result channels
   - Error channels
   - Panic recovery in goroutines
   - Context-based cancellation
   - Deadlock avoidance strategies

6. **Practical Application: Data Pipeline**
   - Real-world example of a data processing pipeline
   - Multi-stage pipeline with generators, processors, and aggregators
   - Worker pool for parallel processing
   - Error handling in a production-like scenario

## Key Takeaways

1. Channels are typed conduits for communication between goroutines
2. Unbuffered channels block until both sender and receiver are ready
3. Buffered channels only block when the buffer is full
4. The select statement allows you to wait on multiple channel operations
5. Proper error handling and timeout mechanisms are essential for robust concurrent programs

## Project Structure

- `main.go` - Main program that runs all examples
- `channelpatterns/` - Directory containing individual pattern demonstrations:
  - `basic.go` - Basic channel operations
  - `buffered.go` - Buffered vs unbuffered channels
  - `sync.go` - Synchronization patterns
  - `select.go` - Select statement examples
  - `errors.go` - Error handling patterns
  - `practical_example.go` - Real-world data pipeline example
  - `learning_mode.go` - Step-by-step explanations for learning mode