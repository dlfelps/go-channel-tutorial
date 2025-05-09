package channelpatterns

import (
	"fmt"
	"time"
)

// DemonstrateBasicChannels shows the basic operation of channels for communication between goroutines
func DemonstrateBasicChannels() {
	fmt.Println("Basic Channel Communication Example")
	fmt.Println("Channels are a typed conduit through which you can send and receive values with the channel operator (<-).")

	// Create an unbuffered channel
	ch := make(chan string)

	// Start a goroutine that sends a message on the channel
	go func() {
		fmt.Println("Goroutine: Starting work...")
		// Simulate some work being done
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine: Work complete, sending result to channel")
		// Send a message on the channel
		ch <- "Hello from goroutine!"
		fmt.Println("Goroutine: Message sent successfully")
	}()

	// Main goroutine
	fmt.Println("Main: Waiting for message from goroutine...")
	// Receive message from the channel - this will block until a message is available
	msg := <-ch
	fmt.Println("Main: Received message:", msg)

	fmt.Println("\nUnidirectional Channels Example")
	fmt.Println("Channels can be restricted to send-only or receive-only.")
	
	// Demonstrate unidirectional channels
	demonstrateUnidirectionalChannels()
	
	fmt.Println("\nIterating over Channels Example")
	fmt.Println("We can use a for-range loop to receive values from a channel until it's closed.")
	
	// Demonstrate channel iteration and closing
	demonstrateChannelIteration()
}

// demonstrateUnidirectionalChannels shows how to use send-only and receive-only channel types
func demonstrateUnidirectionalChannels() {
	// Create a channel
	ch := make(chan int)
	
	// Start a goroutine that sends values
	// The chan<- int parameter means this function can only send to the channel, not receive
	go sender(ch)
	
	// Start a goroutine that receives values
	// The <-chan int parameter means this function can only receive from the channel, not send
	go receiver(ch)
	
	// Let the goroutines run for a moment
	time.Sleep(3 * time.Second)
}

// sender demonstrates a function that uses a send-only channel
func sender(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Sender: Sending %d\n", i)
		ch <- i
		time.Sleep(500 * time.Millisecond)
	}
	// Signal we're done by closing the channel
	close(ch)
	fmt.Println("Sender: Channel closed, no more messages will be sent")
}

// receiver demonstrates a function that uses a receive-only channel
func receiver(ch <-chan int) {
	// Receive until the channel is closed
	for num := range ch {
		fmt.Printf("Receiver: Got %d\n", num)
	}
	fmt.Println("Receiver: Channel is closed, no more messages to receive")
}

// demonstrateChannelIteration shows how to iterate over a channel and handle closing
func demonstrateChannelIteration() {
	// Create a channel
	ch := make(chan string)
	
	// Start a goroutine that sends multiple values and then closes the channel
	go func() {
		messages := []string{"First message", "Second message", "Third message", "Final message"}
		for _, msg := range messages {
			ch <- msg
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Println("Channel Iterator: Closing channel after sending all messages")
		close(ch)
	}()
	
	// Receive messages until the channel is closed
	fmt.Println("Main: Receiving messages using for-range loop:")
	for msg := range ch {
		fmt.Println("  Received:", msg)
	}
	fmt.Println("Main: No more messages in channel (channel closed)")
}
