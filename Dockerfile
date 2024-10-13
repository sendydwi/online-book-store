# Stage 1: Build the application
FROM golang:1.23-alpine3.20 AS builder
WORKDIR /usr/app
COPY . .
RUN go mod download
RUN go build -o bin ./cmd/main.go

# Stage 2: Create a lightweight runtime image
FROM alpine:3.20  AS image
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /usr/app/

COPY --from=builder /usr/app/bin bin
COPY --from=builder /usr/app/.env .env

CMD ["./bin"]