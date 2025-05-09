package channelpatterns

import (
	"fmt"
	"time"
)

// DemonstrateBufferedChannels shows the differences between buffered and unbuffered channels
func DemonstrateBufferedChannels() {
	fmt.Println("Buffered vs Unbuffered Channels")
	fmt.Println("Buffered channels can hold a limited number of values without a receiver being ready.")
	
	// Demonstrate an unbuffered channel
	fmt.Println("\n1. Unbuffered Channel Example:")
	demonstrateUnbufferedChannel()
	
	// Demonstrate a buffered channel
	fmt.Println("\n2. Buffered Channel Example:")
	demonstrateBufferedChannel()
	
	// Demonstrate buffer overflow
	fmt.Println("\n3. Buffer Capacity and Blocking Example:")
	demonstrateBufferCapacity()
}

// demonstrateUnbufferedChannel shows how unbuffered channels block until both sender and receiver are ready
func demonstrateUnbufferedChannel() {
	// Create an unbuffered channel
	unbufferedCh := make(chan string)
	
	fmt.Println("Unbuffered channels block until both sender and receiver are ready")
	
	// Start a goroutine that will send a message
	go func() {
		fmt.Println("Sender: Preparing to send message on unbuffered channel...")
		time.Sleep(2 * time.Second) // Simulate some preparation work
		fmt.Println("Sender: About to send message - will block until receiver is ready")
		unbufferedCh <- "Hello from unbuffered channel!"
		fmt.Println("Sender: Message sent successfully")
	}()
	
	// Wait a moment to let the goroutine start
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Main: Will wait 3 seconds before receiving (sender will be blocked)")
	time.Sleep(3 * time.Second)
	
	// Receive the message
	fmt.Println("Main: Ready to receive message")
	msg := <-unbufferedCh
	fmt.Println("Main: Received message:", msg)
}

// demonstrateBufferedChannel shows how buffered channels allow sending without an immediate receiver
func demonstrateBufferedChannel() {
	// Create a buffered channel with capacity of 2
	bufferedCh := make(chan string, 2)
	
	fmt.Println("Buffered channels can hold a limited number of values without blocking")
	
	// Send messages without a receiver ready
	fmt.Println("Main: Sending first message to buffered channel")
	bufferedCh <- "First message"
	fmt.Println("Main: First message sent (no blocking)")
	
	fmt.Println("Main: Sending second message to buffered channel")
	bufferedCh <- "Second message"
	fmt.Println("Main: Second message sent (no blocking)")
	
	// Now receive the messages
	fmt.Println("Main: Reading first message")
	fmt.Println("  Received:", <-bufferedCh)
	
	fmt.Println("Main: Reading second message")
	fmt.Println("  Received:", <-bufferedCh)
}

// demonstrateBufferCapacity shows what happens when a buffered channel becomes full
func demonstrateBufferCapacity() {
	// Create a buffered channel with capacity of 2
	ch := make(chan int, 2)
	
	fmt.Println("When a buffered channel is full, the sender blocks until space is available")
	fmt.Printf("Creating a channel with buffer capacity of 2\n")
	
	// Start a goroutine that will send more messages than the buffer can hold
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("Sender: Attempting to send value %d\n", i)
			ch <- i
			fmt.Printf("Sender: Sent value %d successfully\n", i)
		}
		fmt.Println("Sender: Closing channel after sending all values")
		close(ch)
	}()
	
	// Receive the messages with deliberate delays to demonstrate blocking
	time.Sleep(1 * time.Second) // Let the sender fill the buffer
	
	fmt.Println("\nBuffer should now be full (2 items). Next send will block until we receive.")
	time.Sleep(2 * time.Second) // Give time to see the sender is blocked
	
	// Receive all messages
	fmt.Println("\nMain: Now receiving all messages")
	for num := range ch {
		fmt.Printf("Main: Received %d\n", num)
		time.Sleep(500 * time.Millisecond) // Slow receiver to demonstrate sender blocking
	}
	
	fmt.Println("Main: All messages received and channel is closed")
}
