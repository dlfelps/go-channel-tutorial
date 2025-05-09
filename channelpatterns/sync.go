package channelpatterns

import (
	"fmt"
	"sync"
	"time"
)

// DemonstrateChannelSync shows various patterns for synchronization using channels
func DemonstrateChannelSync() {
	fmt.Println("Channel Synchronization Patterns")
	fmt.Println("Channels can be used to synchronize execution between goroutines.")
	
	// Demonstrate simple synchronization
	fmt.Println("\n1. Simple Synchronization Example:")
	demonstrateSimpleSync()
	
	// Demonstrate worker pool pattern
	fmt.Println("\n2. Worker Pool Pattern Example:")
	demonstrateWorkerPool()
	
	// Demonstrate fan-out/fan-in pattern
	fmt.Println("\n3. Fan-Out/Fan-In Pattern Example:")
	demonstrateFanOutFanIn()
	
	// Demonstrate using channels as semaphores
	fmt.Println("\n4. Channel as Semaphore Example:")
	demonstrateSemaphore()
}

// demonstrateSimpleSync shows how to wait for a goroutine to finish
func demonstrateSimpleSync() {
	// Create a channel for synchronization
	done := make(chan bool)
	
	fmt.Println("We can use a channel to know when a goroutine completes its work")
	
	// Start a goroutine that simulates work
	go func() {
		fmt.Println("Worker: Starting task...")
		time.Sleep(2 * time.Second) // Simulate work
		fmt.Println("Worker: Task complete")
		
		// Signal that the work is done
		done <- true
	}()
	
	// Wait for the goroutine to finish
	fmt.Println("Main: Waiting for worker to complete...")
	<-done
	fmt.Println("Main: Received completion signal, continuing execution")
}

// demonstrateWorkerPool shows the worker pool pattern using channels
func demonstrateWorkerPool() {
	fmt.Println("The Worker Pool pattern distributes work among multiple goroutines")
	
	const numJobs = 5
	const numWorkers = 3
	
	// Create channels for jobs and results
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	
	// Start workers
	fmt.Printf("Starting %d workers\n", numWorkers)
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}
	
	// Send jobs
	fmt.Println("Sending jobs to workers")
	for j := 1; j <= numJobs; j++ {
		fmt.Printf("  Sending job %d\n", j)
		jobs <- j
	}
	close(jobs)
	fmt.Println("All jobs sent, closed jobs channel")
	
	// Collect results
	fmt.Println("Collecting results:")
	for a := 1; a <= numJobs; a++ {
		result := <-results
		fmt.Printf("  Result: %d\n", result)
	}
}

// worker processes jobs and sends results
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d: started job %d\n", id, j)
		time.Sleep(time.Duration(500*id) * time.Millisecond) // Simulate varying work times
		fmt.Printf("Worker %d: finished job %d\n", id, j)
		results <- j * 2 // Send result (just double the job number for this example)
	}
}

// demonstrateFanOutFanIn shows the fan-out/fan-in pattern
func demonstrateFanOutFanIn() {
	fmt.Println("The Fan-Out/Fan-In pattern distributes work and then collects results")
	
	// Generate some work
	work := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			work <- i
		}
		close(work)
	}()
	
	// Fan-out to 3 workers
	fmt.Println("Fanning out work to 3 workers")
	c1 := fanOut(work)
	c2 := fanOut(work)
	c3 := fanOut(work)
	
	// Fan-in from 3 results channels to 1
	fmt.Println("Fanning in results from workers")
	merged := fanIn(c1, c2, c3)
	
	// Collect all results
	fmt.Println("Results:")
	for result := range merged {
		fmt.Printf("  %d\n", result)
	}
}

// fanOut processes work from a channel and returns a channel of results
func fanOut(work <-chan int) <-chan int {
	resultCh := make(chan int)
	go func() {
		for n := range work {
			// Simulate processing by squaring the input
			time.Sleep(100 * time.Millisecond)
			resultCh <- n * n
		}
		close(resultCh)
	}()
	return resultCh
}

// fanIn merges multiple channels into one channel
func fanIn(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	merged := make(chan int)
	
	// Start a goroutine for each input channel
	output := func(ch <-chan int) {
		for n := range ch {
			merged <- n
		}
		wg.Done()
	}
	
	wg.Add(len(channels))
	for _, ch := range channels {
		go output(ch)
	}
	
	// Start a goroutine to close the merged channel when all input channels are done
	go func() {
		wg.Wait()
		close(merged)
	}()
	
	return merged
}

// demonstrateSemaphore shows how to use channels as semaphores to limit concurrency
func demonstrateSemaphore() {
	fmt.Println("We can use buffered channels as semaphores to limit concurrent operations")
	
	// Create a semaphore channel with capacity 3
	const maxConcurrent = 3
	semaphore := make(chan struct{}, maxConcurrent)
	var wg sync.WaitGroup
	
	// Start many goroutines that all want to run concurrently
	for i := 1; i <= 9; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			fmt.Printf("Task %d: Waiting for semaphore slot\n", id)
			semaphore <- struct{}{} // Acquire semaphore
			fmt.Printf("Task %d: Acquired semaphore slot, starting work\n", id)
			
			// Do some work
			time.Sleep(1 * time.Second)
			
			fmt.Printf("Task %d: Completed work, releasing semaphore slot\n", id)
			<-semaphore // Release semaphore
		}(i)
	}
	
	// Wait for all goroutines to complete
	wg.Wait()
	fmt.Println("All tasks completed")
}
