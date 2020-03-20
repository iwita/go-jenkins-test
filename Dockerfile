FROM golang:latest

LABEL maintainer="Achilleas Tzenetopoulos <atzenetopoulos@microlab.ntua.gr>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY cyos/go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

#Copy the source from the current directory to the Working directory inside the Container
COPY cyos/ .

# Build the Go app
RUN go build -o bin/main cmd/cyosweb/main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./bin/main"]