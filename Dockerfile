# Build image
FROM golang:1.16-alpine AS build

WORKDIR /app
COPY . .
RUN go build -v -o bin/main main.go

# Runtime image
FROM alpine:latest

WORKDIR /app
COPY --from=build /app/bin/main ./
CMD ["./main"]
