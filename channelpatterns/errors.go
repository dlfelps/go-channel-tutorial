package channelpatterns

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// Result represents the result of an operation that may fail
type Result struct {
	Value interface{}
	Err   error
}

// DemonstrateErrorHandling shows patterns for handling errors in concurrent operations
func DemonstrateErrorHandling() {
	fmt.Println("Error Handling Patterns with Channels")
	fmt.Println("Proper error handling is essential for robust concurrent programs.")
	
	// Demonstrate basic error handling with result channels
	fmt.Println("\n1. Result Channels Example:")
	demonstrateResultChannel()
	
	// Demonstrate using a dedicated error channel
	fmt.Println("\n2. Error Channel Example:")
	demonstrateErrorChannel()
	
	// Demonstrate recovering from panics in goroutines
	fmt.Println("\n3. Panic Recovery Example:")
	demonstratePanicRecovery()
	
	// Demonstrate context-based cancellation
	fmt.Println("\n4. Context-Based Cancellation Example:")
	demonstrateChannelCancellation()
	
	// Demonstrate avoiding deadlocks
	fmt.Println("\n5. Deadlock Avoidance Example:")
	demonstrateDeadlockAvoidance()
}

// demonstrateResultChannel shows how to use a Result struct to return errors from goroutines
func demonstrateResultChannel() {
	resultCh := make(chan Result, 3)
	
	// Start several tasks that may succeed or fail
	tasks := []int{1, 0, 2} // Division by zero will cause an error for the second task
	
	for _, denominator := range tasks {
		go func(denom int) {
			fmt.Printf("Task dividing 10 by %d starting...\n", denom)
			time.Sleep(500 * time.Millisecond)
			
			result := Result{}
			if denom == 0 {
				result.Err = errors.New("division by zero")
			} else {
				result.Value = 10 / denom
			}
			
			resultCh <- result
		}(denominator)
	}
	
	// Collect and handle results including errors
	fmt.Println("Collecting results and handling errors:")
	for i := 0; i < len(tasks); i++ {
		res := <-resultCh
		if res.Err != nil {
			fmt.Printf("  Task failed: %v\n", res.Err)
		} else {
			fmt.Printf("  Task succeeded: result = %v\n", res.Value)
		}
	}
}

// demonstrateErrorChannel shows using a dedicated channel for errors
func demonstrateErrorChannel() {
	dataCh := make(chan int)
	errCh := make(chan error)
	
	// Start a task that might fail
	go func() {
		fmt.Println("Starting task that might fail...")
		time.Sleep(1 * time.Second)
		
		// Generate a random error condition
		if time.Now().UnixNano()%2 == 0 {
			errCh <- errors.New("something went wrong")
			return
		}
		
		// Task succeeded
		dataCh <- 42
	}()
	
	// Wait for either data or an error
	fmt.Println("Waiting for task to complete or fail...")
	select {
	case data := <-dataCh:
		fmt.Println("Task completed successfully with result:", data)
	case err := <-errCh:
		fmt.Println("Task failed with error:", err)
	}
}

// demonstratePanicRecovery shows how to recover from panics in goroutines
func demonstratePanicRecovery() {
	// Channel to signal task completion or failure
	done := make(chan bool)
	
	// Start a task that will panic
	go func() {
		// Set up deferred panic recovery
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered from panic: %v\n", r)
				done <- false // Signal that the task failed
			}
		}()
		
		fmt.Println("Starting task that will panic...")
		time.Sleep(500 * time.Millisecond)
		
		// This will cause a panic
		fmt.Println("About to panic...")
		panic("deliberate panic for demonstration")
		
		// This code will not be reached
		done <- true
	}()
	
	// Wait for task completion
	if <-done {
		fmt.Println("Task completed successfully")
	} else {
		fmt.Println("Task failed but was recovered")
	}
}

// demonstrateChannelCancellation shows how to cancel operations using channels
func demonstrateChannelCancellation() {
	// Channel to signal cancellation
	cancel := make(chan struct{})
	
	// Channel to report completion
	done := make(chan struct{})
	
	// Start a task that can be cancelled
	go func() {
		fmt.Println("Starting long-running task...")
		
		// Simulate work that periodically checks for cancellation
		for i := 1; i <= 5; i++ {
			select {
			case <-cancel:
				fmt.Println("Task cancelled, cleaning up and exiting")
				done <- struct{}{}
				return
			case <-time.After(300 * time.Millisecond):
				fmt.Printf("Task working... step %d/5\n", i)
			}
		}
		
		fmt.Println("Task completed successfully")
		done <- struct{}{}
	}()
	
	// Simulate waiting and then deciding to cancel
	fmt.Println("Will cancel the task after 1 second")
	time.Sleep(1 * time.Second)
	
	fmt.Println("Sending cancellation signal")
	close(cancel) // Signal cancellation by closing the channel
	
	// Wait for the task to acknowledge cancellation
	<-done
	fmt.Println("Main: received confirmation that task has stopped")
}

// demonstrateDeadlockAvoidance shows how to avoid deadlocks in channel operations
func demonstrateDeadlockAvoidance() {
	fmt.Println("A deadlock occurs when all goroutines are blocked waiting on channels")
	fmt.Println("Here are some strategies to avoid deadlocks:")
	
	// Demonstrate the buffer strategy
	fmt.Println("\n  Strategy 1: Use buffered channels when appropriate")
	demonstrateBufferStrategy()
	
	// Demonstrate the select default case strategy
	fmt.Println("\n  Strategy 2: Use select with default case for non-blocking operations")
	demonstrateSelectStrategy()
	
	// Demonstrate timeout strategy
	fmt.Println("\n  Strategy 3: Use timeouts to prevent indefinite blocking")
	demonstrateTimeoutStrategy()
	
	// Demonstrate proper channel closure
	fmt.Println("\n  Strategy 4: Always ensure channels are properly closed")
	demonstrateCloseStrategy()
}

// demonstrateBufferStrategy shows using buffered channels to avoid deadlocks
func demonstrateBufferStrategy() {
	// This could deadlock if unbuffered
	ch := make(chan int, 1)
	
	fmt.Println("  Using a buffered channel to avoid deadlock in a single goroutine")
	
	// This won't block because the channel is buffered
	ch <- 1
	fmt.Println("  Sent value to buffered channel (no deadlock)")
	
	// Now we can receive it
	val := <-ch
	fmt.Println("  Received value:", val)
}

// demonstrateSelectStrategy shows using select to avoid blocking indefinitely
func demonstrateSelectStrategy() {
	ch := make(chan int)
	
	fmt.Println("  Using select with default to make non-blocking channel operations")
	
	// Try to receive without blocking
	select {
	case val := <-ch:
		fmt.Println("  Received value:", val)
	default:
		fmt.Println("  Would block on receive, taking alternative action")
	}
	
	// Try to send without blocking
	select {
	case ch <- 1:
		fmt.Println("  Sent value to channel")
	default:
		fmt.Println("  Would block on send, taking alternative action")
	}
}

// demonstrateTimeoutStrategy shows using timeouts to avoid blocking forever
func demonstrateTimeoutStrategy() {
	ch := make(chan int)
	
	fmt.Println("  Using a timeout to avoid waiting forever")
	
	// Try to receive with a timeout
	select {
	case val := <-ch:
		fmt.Println("  Received value:", val)
	case <-time.After(1 * time.Second):
		fmt.Println("  Timed out waiting for value, continuing execution")
	}
}

// demonstrateCloseStrategy shows proper channel closing to avoid deadlocks
func demonstrateCloseStrategy() {
	ch := make(chan int)
	var wg sync.WaitGroup
	
	// Sender
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch) // Important: close the channel when done sending
		
		for i := 1; i <= 3; i++ {
			ch <- i
		}
		fmt.Println("  Sender: Finished sending all values, closed channel")
	}()
	
	// Receiver
	wg.Add(1)
	go func() {
		defer wg.Done()
		
		// This loop will exit when the channel is closed
		for val := range ch {
			fmt.Printf("  Receiver: Got value %d\n", val)
		}
		fmt.Println("  Receiver: Channel closed, exiting loop safely")
	}()
	
	// Wait for both goroutines to finish
	wg.Wait()
	fmt.Println("  Both sender and receiver completed without deadlock")
}
