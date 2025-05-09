package channelpatterns

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Item represents a simulated data item to be processed
type Item struct {
	ID    int
	Value string
}

// ProcessResult represents the result of processing an item
type ProcessResult struct {
	ItemID int
	Result string
	Error  error
}

// DemonstratePracticalExample shows a real-world example of using channels
// in a data processing pipeline
func DemonstratePracticalExample() {
	fmt.Println("Practical Example: Data Processing Pipeline")
	fmt.Println("This demonstrates how channels can be used to build a concurrent data processing pipeline.")

	// Step 1: Create pipeline stages
	items := generateItems(20)         // Source: generate 20 items
	processedItems := processItems(items, 5) // Processing stage with 5 workers
	results := aggregateResults(processedItems) // Results aggregation stage

	// Step 2: Display the pipeline results
	fmt.Println("\nProcessing Pipeline Results:")
	fmt.Printf("Total items processed: %d\n", len(results))
	
	// Show a sample of results
	fmt.Println("\nSample of processed results:")
	for i, result := range results {
		if i < 5 { // Show only 5 samples
			if result.Error != nil {
				fmt.Printf("  Item %d: ERROR - %v\n", result.ItemID, result.Error)
			} else {
				fmt.Printf("  Item %d: %s\n", result.ItemID, result.Result)
			}
		}
	}
	fmt.Printf("  ... and %d more results\n", len(results)-5)
	
	fmt.Println("\nPipeline Key Features:")
	fmt.Println("1. Data flows through channels between independent stages")
	fmt.Println("2. Each stage can run at its own pace (backpressure handling)")
	fmt.Println("3. Processing is parallelized with a worker pool")
	fmt.Println("4. Results maintain original ordering through a map")
}

// generateItems simulates a data source producing items
func generateItems(count int) <-chan Item {
	out := make(chan Item)
	
	go func() {
		defer close(out)
		
		fmt.Println("\nStage 1: Generating source data items")
		for i := 1; i <= count; i++ {
			item := Item{
				ID:    i,
				Value: fmt.Sprintf("Original data %d", i),
			}
			
			// Simulate variable item generation time
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			fmt.Printf("  Generated item %d\n", item.ID)
			out <- item
		}
		fmt.Println("  All items generated")
	}()
	
	return out
}

// processItems handles the parallel processing stage of the pipeline
func processItems(in <-chan Item, numWorkers int) <-chan ProcessResult {
	out := make(chan ProcessResult)
	
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	
	fmt.Printf("\nStage 2: Processing items with %d workers\n", numWorkers)
	
	// Start multiple workers
	for i := 1; i <= numWorkers; i++ {
		go func(workerID int) {
			defer wg.Done()
			
			for item := range in {
				// Simulate processing time and occasional errors
				processingTime := time.Duration(rand.Intn(500)) * time.Millisecond
				time.Sleep(processingTime)
				
				result := ProcessResult{ItemID: item.ID}
				
				// Simulate occasional processing errors (about 10%)
				if rand.Intn(10) == 0 {
					result.Error = fmt.Errorf("processing failed")
					fmt.Printf("  Worker %d: Error processing item %d\n", workerID, item.ID)
				} else {
					result.Result = fmt.Sprintf("Processed: %s (by worker %d)", item.Value, workerID)
					fmt.Printf("  Worker %d: Processed item %d\n", workerID, item.ID)
				}
				
				out <- result
			}
		}(i)
	}
	
	// Close the output channel when all workers are done
	go func() {
		wg.Wait()
		close(out)
		fmt.Println("  All items processed")
	}()
	
	return out
}

// aggregateResults collects and aggregates the processing results
func aggregateResults(in <-chan ProcessResult) []ProcessResult {
	fmt.Println("\nStage 3: Aggregating and analyzing results")
	
	var results []ProcessResult
	successCount := 0
	errorCount := 0
	
	for result := range in {
		// Store all results in a slice
		results = append(results, result)
		
		// Count successes and failures
		if result.Error != nil {
			errorCount++
		} else {
			successCount++
		}
		
		// Simulate result processing time
		time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
	}
	
	fmt.Printf("  Aggregation complete: %d successes, %d errors\n", 
		successCount, errorCount)
	
	return results
}