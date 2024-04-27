# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory in the container
WORKDIR /go/src/go-commander

# Copy the local package files to the container's workspace
COPY . .

# Download all dependencies
RUN go mod download

# Build the Go app
RUN go build -o main .

# Create a non-root user
RUN useradd -u 10001 -U -m appuser

# Change ownership of the working directory to the non-root user
RUN chown -R appuser:appuser /go/src/go-commander

# Set the user for subsequent commands
USER appuser

# Open a bash shell
CMD ["./main"]
