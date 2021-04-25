# Build image
##################################################
FROM golang:1.16-alpine AS build

WORKDIR /app

# Download dependencies and cache them
COPY go.* ./
RUN go mod download

# Build the program
COPY . .
RUN go build -v -o bin/bot ./...

# Runtime image
##################################################
FROM alpine:latest

WORKDIR /app
COPY --from=build /app/bin/bot ./
CMD ["./bot"]
