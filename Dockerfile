# Use the official Golang image for building and running the Go application
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the entire project directory to the working directory
COPY . .

# Run go mod tidy to ensure dependencies are installed
RUN go mod tidy

# Build the Go application
RUN go build -o main ./cmd/http/main.go

# Expose the port the application will run on
EXPOSE 8080

# Set the command to run the application
CMD ["./main"]
