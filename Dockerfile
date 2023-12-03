# Use an official Go runtime as the parent image
FROM golang:latest as builder

# Set the working directory in the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Download the Go modules
RUN go mod download

# Build the application
RUN make build

# Expose port 8000 for the app
EXPOSE 8000

# Copy configuration file
COPY config.local.yaml /app/config/config.local.yaml

# Run the binary on container startup
CMD ["./build/user_account", "server"]