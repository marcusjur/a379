FROM golang:1.23-alpine

# Set the working directory
WORKDIR /app-backend

# Copy the Go mod
COPY go.mod go.sum ./

# Download the dependencies
RUN apk add --no-cache gcc musl-dev
RUN go mod download

# Copy the Go application files
COPY main.go ./
COPY session/session.go ./session/

# Build the Go app
RUN CGO_ENABLED=1 go build -o libsession.so -buildmode=c-shared ./session/session.go
RUN go build -o main .

# Expose the ports
EXPOSE 8080
EXPOSE 3306

# Run the application
CMD ["./main"]
