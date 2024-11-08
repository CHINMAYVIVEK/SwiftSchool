# Step 1: Use an official Go image to build the application
FROM golang:1.20 AS builder

# Step 2: Set the Current Working Directory inside the container
WORKDIR /app

# Step 3: Copy go mod and sum files
COPY go.mod go.sum ./

# Step 4: Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Step 5: Copy the source code into the container
COPY . .

# Step 6: Build the Go app
RUN go build -o main .

# Step 7: Use a smaller base image for the final image
FROM alpine:latest  

# Step 8: Install necessary dependencies in the final image
RUN apk --no-cache add ca-certificates

# Step 9: Set the Current Working Directory inside the container
WORKDIR /root/

# Step 10: Copy the binary from the builder image
COPY --from=builder /app/main .

# Step 11: Expose the port the app runs on
EXPOSE 8080

# Step 12: Command to run the application
CMD ["./main"]
