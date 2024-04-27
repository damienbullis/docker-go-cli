# Variables
IMAGE_NAME := go_commander
CONTAINER_NAME := go_commander_container

# Default target
.PHONY: build run clean test format
.DEFAULT_GOAL := run

# Build the Docker image
build:
	docker build -t $(IMAGE_NAME) .

# Run the Docker container
run:
	docker run --rm -it --name $(CONTAINER_NAME) $(IMAGE_NAME)

# Clean up
clean:
	docker stop $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)

# Run tests
test:
	docker run --rm $(IMAGE_NAME) go test ./...

# Format code
format:
	docker run --rm $(IMAGE_NAME) go fmt ./...
