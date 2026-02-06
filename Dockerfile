# Step 1: Use the official Go image to BUILD the app
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o health_monitor health_monitor.go

# Step 2: Use a tiny "Alpine" image to RUN the app
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/health_monitor .

# Command to run the executable
ENTRYPOINT ["./health_monitor"]