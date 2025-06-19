# ChronoCashe
ChronoCashe is a lightweight, time-windowed in-memory key-value cache built with Go. It allows you to store key-value pairs that are only accessible within a specified time window, making it ideal for use cases like temporary data storage, rate limiting, or time-sensitive promotions.
Features

- Time-based Availability: Store keys with a defined availability window using available_from and available_until timestamps.
- HTTP API: Simple RESTful API for creating, retrieving, deleting, and listing cache entries.
- Thread Safety: Concurrent access handled safely with read-write locks.
- Periodic Cleanup: Automatically removes expired keys via a background scheduler.
- Minimal Dependencies: Built with the Go standard library and the Chi router for a lightweight footprint.

## Prerequisites

Go 1.24.4 or later
Git (for cloning the repository)

## Installation

### Clone the repository:
```bash
git clone https://github.com/yourusername/chronocashe.git
cd chronocashe
```


### Install dependencies:
```bash
go mod tidy
```


### Run the server:
```bash
go run main.go
```

The server will start on http://localhost:8080.


### Time Format
The available_from and available_until fields must be in RFC3339 format (e.g., 2025-06-19T15:00:00Z).
## Usage Examples
### Set a Key
```bash
curl -X PUT http://localhost:8080/cache/test \
  -H "Content-Type: application/json" \
  -d '{"value":"hello","available_from":"2025-06-19T15:00:00Z","available_until":"2025-06-19T16:00:00Z"}'
```

Response:
```bash
{
  "key": "test",
  "value": "hello",
  "available_from": "2025-06-19T15:00:00Z",
  "available_until": "2025-06-19T16:00:00Z"
}
```

### Get a Key
```bash
curl http://localhost:8080/cache/test
```

Response:
```bash
{
  "key": "test",
  "value": "hello"
}
```

### Delete a Key
```bash
curl -X DELETE http://localhost:8080/cache/test
```

Response: (No content, status 204)

### List Active Keys
```bash
curl http://localhost:8080/cache
```

Response:
```bash
[
  {
    "key": "test",
    "value": "hello",
    "available_from": "2025-06-19T15:00:00Z",
    "available_until": "2025-06-19T16:00:00Z"
  }
]
```

## Project Structure
chronocashe/

├── internal/

│   ├── api/          # HTTP API handlers

│   ├── cache/        # Cache logic and engine

│   ├── models/       # Data structures

│   ├── scheduler/    # Background cleanup scheduler

├── go.mod            # Go module definition

├── go.sum            # Dependency checksums

├── main.go           # Application entry point

└── README.md         # Project documentation

## Contributing
Contributions are welcome! Please follow these steps:

### Fork the repository.
- Create a new branch (git checkout -b feature/your-feature).
- Make your changes and commit them (git commit -m "Add your feature").
- Push to your branch (git push origin feature/your-feature).
- Open a pull request.

### Development Setup
To run tests (once added):
go test ./...

## License
This project is licensed under the MIT License. See the LICENSE file for details.

## Contact
For questions or feedback, please open an issue on the GitHub repository.
