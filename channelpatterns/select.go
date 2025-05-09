package channelpatterns

import (
	"fmt"
	"math/rand"
	"time"
)

// DemonstrateSelect shows how to use the select statement to handle multiple channels
func DemonstrateSelect() {
	fmt.Println("Select Statement with Channels")
	fmt.Println("The select statement lets you wait on multiple channel operations simultaneously.")
	
	// Demonstrate basic select usage
	fmt.Println("\n1. Basic Select Example:")
	demonstrateBasicSelect()
	
	// Demonstrate using a timeout with select
	fmt.Println("\n2. Timeout Example:")
	demonstrateSelectTimeout()
	
	// Demonstrate using default case in select
	fmt.Println("\n3. Default Case Example (Non-blocking check):")
	demonstrateSelectDefault()
	
	// Demonstrate using select for multiplexing
	fmt.Println("\n4. Multiplexing Example:")
	demonstrateMultiplexing()
	
	// Demonstrate channel closing detection with select
	fmt.Println("\n5. Channel Closing Detection Example:")
	demonstrateSelectClosing()
}

// demonstrateBasicSelect shows the basic usage of select to wait on multiple channels
func demonstrateBasicSelect() {
	c1 := make(chan string)
	c2 := make(chan string)
	
	// Start two goroutines, each sending data on a different channel
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "Message from channel 1"
	}()
	
	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "Message from channel 2"
	}()
	
	// Wait for messages from both channels using select
	fmt.Println("Waiting for messages from two channels...")
	
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received:", msg1)
		case msg2 := <-c2:
			fmt.Println("Received:", msg2)
		}
	}
}

// demonstrateSelectTimeout shows how to implement a timeout using select
func demonstrateSelectTimeout() {
	c := make(chan string)
	
	// Start a goroutine that may or may not send a value
	go func() {
		// Simulate either a quick or slow response
		rand.Seed(time.Now().UnixNano())
		delay := rand.Intn(3) + 1 // 1, 2, or 3 seconds
		fmt.Printf("Operation will take %d seconds...\n", delay)
		time.Sleep(time.Duration(delay) * time.Second)
		c <- "Operation complete"
	}()
	
	// Wait for the operation with a timeout
	fmt.Println("Waiting for operation with a 2-second timeout...")
	
	select {
	case result := <-c:
		fmt.Println("Success:", result)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout: Operation took too long")
	}
	
	// We should wait a moment to allow any pending goroutine to complete
	// to avoid "goroutine leak" where a goroutine is blocked trying to send on a channel
	time.Sleep(3 * time.Second)
}

// demonstrateSelectDefault shows how to use the default case for non-blocking operations
func demonstrateSelectDefault() {
	messages := make(chan string)
	signals := make(chan bool)
	
	// Try a non-blocking receive
	fmt.Println("Non-blocking receive operation:")
	select {
	case msg := <-messages:
		fmt.Println("Received message:", msg)
	default:
		fmt.Println("No message received (would block)")
	}
	
	// Try a non-blocking send
	msg := "Hello"
	fmt.Println("Non-blocking send operation:")
	select {
	case messages <- msg:
		fmt.Println("Sent message:", msg)
	default:
		fmt.Println("No message sent (would block)")
	}
	
	// Try a multi-way non-blocking select
	fmt.Println("Multi-way non-blocking select:")
	select {
	case msg := <-messages:
		fmt.Println("Received message:", msg)
	case sig := <-signals:
		fmt.Println("Received signal:", sig)
	default:
		fmt.Println("No activity (would block)")
	}
}

// demonstrateMultiplexing shows how to use select to combine messages from multiple channels
func demonstrateMultiplexing() {
	c1 := make(chan string)
	c2 := make(chan string)
	
	// Send messages on the first channel periodically
	go func() {
		for i := 1; i <= 3; i++ {
			time.Sleep(200 * time.Millisecond)
			c1 <- fmt.Sprintf("Message %d from source 1", i)
		}
	}()
	
	// Send messages on the second channel periodically
	go func() {
		for i := 1; i <= 3; i++ {
			time.Sleep(300 * time.Millisecond)
			c2 <- fmt.Sprintf("Message %d from source 2", i)
		}
	}()
	
	// Multiplex messages from both channels for a while
	fmt.Println("Multiplexing messages from two sources:")
	
	// Create a timeout for the entire operation
	timeout := time.After(1 * time.Second)
	
	for {
		select {
		case msg1 := <-c1:
			fmt.Println("Source 1:", msg1)
		case msg2 := <-c2:
			fmt.Println("Source 2:", msg2)
		case <-timeout:
			fmt.Println("Timeout reached, stopping multiplexer")
			return
		}
	}
}

// demonstrateSelectClosing shows how to detect when a channel is closed using select
func demonstrateSelectClosing() {
	// Create a channel that will be closed
	ch := make(chan int)
	
	// Start a goroutine that sends values and then closes the channel
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Println("Sender: Closing the channel")
		close(ch)
	}()
	
	// Receive values, detect when the channel is closed
	fmt.Println("Receiving values and detecting channel close:")
	
	for {
		select {
		case val, ok := <-ch:
			if !ok {
				fmt.Println("Channel is closed, exiting loop")
				return
			}
			fmt.Println("Received value:", val)
		}
	}
}
