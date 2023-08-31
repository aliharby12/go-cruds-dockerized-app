# Use the official Go base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container's workspace
COPY . .

# Build the Go application
RUN go build -o main .

# Run air
CMD ["air"]