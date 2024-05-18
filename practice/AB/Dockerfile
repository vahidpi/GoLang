# Use an official Go runtime as a parent image
FROM golang:1.16.10-alpine3.14

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Install MySQL client library
RUN apk add --no-cache mysql-client

# Download and install any required dependencies
RUN go mod download

# Build the Go application
RUN go build -o ascii-service .

# Expose port 8080
EXPOSE 8080

# Define the environment variables
ENV DB_HOST=localhost \
    DB_PORT=3306 \
    DB_USER=root \
    DB_PASSWORD= \
    DB_NAME=ascii

# Run the application
CMD ["./ascii-service"]