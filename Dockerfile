# Stage 1: Build the application
FROM golang:1.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o bookstore

# Stage 2: Create a lightweight runtime image
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/bookstore .
CMD ["./myapp"]
