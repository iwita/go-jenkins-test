FROM golang:latest as builder

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

#### Start a new stage from scratch ####
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from-builder /app/bin/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./bin/main"]
