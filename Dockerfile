# Use an official Golang runtime as a parent image
FROM golang:1.21-alpine

# Set the working directory to /app
WORKDIR /app

# Copy the main.go file from /cmd/api to /app in the container
COPY ./cmd/api/main.go /app
COPY ./go.mod /app
COPY ./go.sum /app
COPY . .

# Build the Go app
RUN apk add --no-cache git \
    && go get -d -v \
    && go install -v \
    && go build -o /deres

EXPOSE 8080

ENTRYPOINT ["/deres"]