# Use an official Golang runtime as a parent image
FROM golang:1.21-alpine

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the Go app
RUN apk add --no-cache git \
    && go get -d -v \
    && go install -v \
    && go build -o /deres

# Expose port 8080 for the container
EXPOSE 8080

# Run the Go app when the container starts
CMD ["/deres"]