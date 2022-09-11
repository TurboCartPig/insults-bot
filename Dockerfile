# Build image
##################################################
FROM golang:1.19.1-alpine AS build

WORKDIR /app

# Download dependencies and cache them
COPY go.* ./
RUN go mod download

# Build the program
COPY . .
RUN go build \
	-v \
	-o bin/bot \
	./...

# Runtime image
##################################################
FROM alpine:3.13

WORKDIR /app
COPY --from=build /app/bin/bot ./
CMD ["./bot"]
