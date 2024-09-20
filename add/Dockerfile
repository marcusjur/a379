FROM golang:1.23-alpine

# Set the working directory
WORKDIR /app-backend

# Copy the Go mod
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the Go application files
COPY main.go ./

# Build the Go app
RUN go build -o main .

# Expose the ports
EXPOSE 8080
EXPOSE 3306

# Run the application
CMD ["./main"]
