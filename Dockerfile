# Use the official Go image as the base image for the build stage
FROM golang:latest AS builder

# Set the working directory in the container
WORKDIR /app

COPY . .

# Remove any previously initialized go.mod and go.sum files
RUN rm -f go.mod go.sum

# Initialize Go modules
RUN go mod init lotery_viking

# Fetch dependencies
RUN go mod tidy

# Build the Go binary
RUN CGO_ENABLED=0 go build -o main cmd/api/main.go

# Use a minimal base image for the production stage
FROM alpine:latest

# Copy the binary from the build stage to the production stage
COPY --from=builder /app/main /main

# Expose the port that the application listens on
EXPOSE 3000

# CMD ["/bin/sh", "-c", "./main serve"]
CMD ["/bin/sh", "-c", "./main init && ./main drop && ./main migrate && ./main seed && ./main serve"]
