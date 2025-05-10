.PHONY: build run list clean

# Binary name
BINARY_NAME=channel-patterns

# Build the go application
build:
	go build -o $(BINARY_NAME) -v

# Run all demonstrations
run:
	go run main.go

# Run the list command to show all learning topics
list:
	go run main.go -list

# Run a specific learning topic
learn:
	@if [ -z "$(TOPIC)" ]; then \
		echo "Please specify a topic with TOPIC=<topic>"; \
		echo "Available topics:"; \
		go run main.go -list; \
	else \
		go run main.go -learn $(TOPIC); \
	fi

# Clean up binary
clean:
	go clean
	rm -f $(BINARY_NAME)

# Run tests
test:
	go test -v ./...