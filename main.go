package main

import (
        "flag"
        "fmt"
        "os"
        "strings"
        "time"

        "channelexamples/channelpatterns"
)

func main() {
        // Define command line flags and parse them right away
        learnMode := flag.String("learn", "", "Enable learning mode for a specific topic (basic, buffered, sync, select, errors, pipeline)")
        showTopics := flag.Bool("list", false, "List all available topics for learning mode")
        
        // Parse the command line arguments
        flag.Parse()
        
        // Check if we should just show the list of topics
        if *showTopics {
                fmt.Println("==== Available Learning Mode Topics ====")
                fmt.Println("Use -learn flag followed by one of these topics:")
                fmt.Println("  basic     - Basic channel operations")
                fmt.Println("  buffered  - Buffered vs unbuffered channels")
                fmt.Println("  sync      - Channel synchronization patterns")
                fmt.Println("  select    - Select statement patterns")
                fmt.Println("  errors    - Error handling with channels")
                fmt.Println("  pipeline  - Data processing pipeline pattern")
                fmt.Println("\nExample: go run main.go -learn basic")
                return
        }
        
        // Check if learning mode was requested
        if *learnMode != "" {
                runLearningMode(*learnMode)
                return
        }
        
        // If no special flags, run all demonstrations
        runAllDemonstrations()
}

// listAvailableTopics displays all available topics for learning mode
func listAvailableTopics() {
        fmt.Println("==== Available Learning Mode Topics ====")
        fmt.Println("Use -learn flag followed by one of these topics:")
        fmt.Println("  basic     - Basic channel operations")
        fmt.Println("  buffered  - Buffered vs unbuffered channels")
        fmt.Println("  sync      - Channel synchronization patterns")
        fmt.Println("  select    - Select statement patterns")
        fmt.Println("  errors    - Error handling with channels")
        fmt.Println("  pipeline  - Data processing pipeline pattern")
        fmt.Println("\nExample: go run main.go -learn basic")
}

// runLearningMode runs the learning mode for a specific topic
func runLearningMode(topic string) {
        // Map of valid topics
        validTopics := map[string]struct{}{
                "basic":    {},
                "buffered": {},
                "sync":     {},
                "select":   {},
                "errors":   {},
                "pipeline": {},
        }

        // Check if topic is valid
        if _, exists := validTopics[topic]; !exists {
                fmt.Printf("Error: Unknown topic '%s'\n", topic)
                listAvailableTopics()
                os.Exit(1)
        }

        fmt.Printf("==== Go Channel Patterns - Learning Mode: %s ====\n", strings.ToUpper(topic))
        fmt.Println("This mode provides detailed step-by-step explanations before the demonstration.")

        // Get the demonstration function for this topic
        var demoFn func()
        switch topic {
        case "basic":
                demoFn = channelpatterns.DemonstrateBasicChannels
        case "buffered":
                demoFn = channelpatterns.DemonstrateBufferedChannels
        case "sync":
                demoFn = channelpatterns.DemonstrateChannelSync
        case "select":
                demoFn = channelpatterns.DemonstrateSelect
        case "errors":
                demoFn = channelpatterns.DemonstrateErrorHandling
        case "pipeline":
                demoFn = channelpatterns.DemonstratePracticalExample
        }

        // Run the learning mode for this topic
        channelpatterns.DemonstrateLearningMode(topic)
        
        // Now run the actual demonstration
        fmt.Println("\n==== LIVE DEMONSTRATION ====")
        demoFn()
        
        fmt.Println("\n==== Learning Session Complete ====")
        fmt.Println("You've learned about " + topic + " channel patterns with detailed explanations and a live demonstration.")
        fmt.Println("Try another topic with: go run main.go -learn <topic>")
        fmt.Println("Or run all demonstrations with: go run main.go")
}

// runAllDemonstrations runs all the channel pattern demonstrations
func runAllDemonstrations() {
        fmt.Println("==== Go Channel Patterns Demonstration ====")
        fmt.Println("This program demonstrates various channel patterns and techniques for effective concurrency in Go.")
        fmt.Println("For detailed explanations, use learning mode with: go run main.go -learn <topic>")
        
        // Define all demonstrations
        examples := []struct {
                name string
                fn   func()
        }{
                {"Basic Channel Communication", channelpatterns.DemonstrateBasicChannels},
                {"Buffered vs Unbuffered Channels", channelpatterns.DemonstrateBufferedChannels},
                {"Channel Synchronization", channelpatterns.DemonstrateChannelSync},
                {"Select Statement", channelpatterns.DemonstrateSelect},
                {"Error Handling", channelpatterns.DemonstrateErrorHandling},
                {"Practical Application: Data Pipeline", channelpatterns.DemonstratePracticalExample},
        }

        // Run each demonstration
        for i, example := range examples {
                fmt.Printf("\n\n%d. %s\n", i+1, example.name)
                fmt.Println(strings.Repeat("-", len(example.name)+4))
                example.fn()
                
                // Add a short pause between examples
                if i < len(examples)-1 {
                        fmt.Println("\nContinuing to the next example in 2 seconds...")
                        time.Sleep(2 * time.Second)
                }
        }
        
        fmt.Println("\n==== Demonstration Complete ====")
        fmt.Println("You've seen various channel patterns that enable effective concurrency in Go.")
        fmt.Println("Remember these key takeaways:")
        fmt.Println("1. Channels are typed conduits for communication between goroutines")
        fmt.Println("2. Unbuffered channels block until both sender and receiver are ready")
        fmt.Println("3. Buffered channels only block when the buffer is full")
        fmt.Println("4. The select statement allows you to wait on multiple channel operations")
        fmt.Println("5. Proper error handling and timeout mechanisms are essential for robust concurrent programs")
        fmt.Println("\nFor detailed explanations, try the learning mode: go run main.go -learn <topic>")
}
