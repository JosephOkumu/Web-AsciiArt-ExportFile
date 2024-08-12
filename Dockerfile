# Use the official Golang image to create a build stage
FROM golang:1.22.5 AS builder

# Create a filesystem in the image for the container
WORKDIR /app

# Copy go mod and sum files first to leverage Docker cache
COPY go.mod ./

# Download dependencies. Dependencies are cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container file system.
COPY . .

# Initialize the go module based on the source code in the container
RUN go mod tidy

# Build the Go app
RUN go build -o main .

# Use a minimal base image for the final stage
FROM debian:bookworm-slim

# Set the current Working Directory inside the container
WORKDIR /app

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Optionally copy static files or other resources if needed
COPY static /app/static
COPY templates /app/templates
COPY web /app/web
COPY banners /app/banners

# Expose port 8080 to the outside world.
EXPOSE 8080

# Apply metadata
LABEL org.opencontainers.image.title="Go Web Server" \
      org.opencontainers.image.description="Website that displays Ascii art" \
      org.opencontainers.image.version="1.0" \
      org.opencontainers.image.author="josotieno and kada"

# Run the executable
CMD ["./main"]
