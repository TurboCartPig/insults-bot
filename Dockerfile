# Build image
##################################################
FROM golang:1.20-alpine3.16 AS build

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
FROM alpine:3.16

WORKDIR /app
COPY --from=build /app/bin/bot ./
CMD ["./bot"]
