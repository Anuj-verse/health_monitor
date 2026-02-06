# Go-Health-Monitor 🚀

A high-performance, concurrent CLI tool built in Go to monitor website availability. This project demonstrates core Go concepts like Goroutines, Channels, and Structs, packaged as a lightweight Docker container.

## 🛠 Features
- **Concurrent Execution**: Uses Goroutines to check multiple URLs simultaneously rather than one-by-one.
- **Thread-Safe Communication**: Implements Go Channels to synchronize results back to the main thread.
- **Robust Error Handling**: Captures network timeouts and DNS errors using a custom `res` struct.
- **Containerized**: Multi-stage Docker build resulting in a tiny footprint (<20MB).

## 🚀 Getting Started

### Prerequisites
- Go 1.21+ (if running locally)
- Docker (if running as a container)

### Running Locally
```bash
go run health_monitor.go
