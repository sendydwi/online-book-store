# Stage 1: Build the application
FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o bin ./cmd/main.go
RUN chmod +x bin
RUN echo $(ls -l)


CMD ["./bin"]

# # Stage 2: Create a lightweight runtime image
# FROM alpine:latest
# WORKDIR /app
# COPY --from=builder /app/bin bin
# COPY --from=builder /app/.env .env

