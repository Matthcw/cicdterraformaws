# Use an official Go runtime as a parent image
FROM golang:1.18-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the current directory contents into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Make port 8080 available to the outside world
EXPOSE 8080

# Run the binary
CMD ["./main"]
