# Use the official Go base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container's workspace
COPY . .

# Build the Go application
RUN go build -o main .

# Install swag
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Generate swag documentation
RUN swag init

# Run migrations
RUN go run migrations/migrate.go

# Install air
RUN go get -u github.com/cosmtrek/air

# Run air
CMD ["air"]