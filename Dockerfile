# Use an official Golang runtime as a parent image
FROM golang:latest

# Ensure color terminal output
ENV TERM=xterm-256color

# Set the working directory in the container
WORKDIR /go/src/go-commander

# Copy the local package files to the container's workspace
COPY . .

# Download all dependencies
RUN go mod download

# Create a non-root user
RUN useradd -u 10001 -U -m appuser

# Change ownership of the working directory and mod cache directory to the non-root user
RUN chown -R appuser:appuser /go/src/go-commander

# Set proper permissions for the go packages
RUN chmod -R 777 /go/pkg

# Set the user for subsequent commands
USER appuser
