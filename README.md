# Go Book API

A simple RESTful API for managing books, built with Go, Gin, and GORM.

## Features

- CRUD operations for books
- PostgreSQL database integration
- Dockerized for easy deployment
- Automated tests with SQLite

## Getting Started

### Prerequisites

- Go 1.25+
- Docker & Docker Compose

### Setup

1. Clone the repository
2. Copy `.env.example` to `.env` and update your database credentials
3. Build and run with Docker Compose:
   ```sh
   docker-compose up --build
   ```
4. The API will be available at `http://localhost:8040`

### API Endpoints

- `POST   /books` - Create a new book
- `GET    /books` - List all books
- `GET    /books/:id` - Get a book by ID
- `PUT    /books/:id` - Update a book
- `DELETE /books/:id` - Delete a book

## Testing

Run tests locally:

```sh
go test ./...
```

Or as part of the Docker build (see Dockerfile).

## Project Structure

- `api/`
  - `handlers/` - HTTP handlers (controllers)
  - `repositories/` - Database access logic (repository pattern)
  - `model.go` - Core Book model
  - `dto.go` - Data Transfer Objects (request/response structs)
  - `response/` - Response helpers (JsonResponse, ResponseJson)
  - `router.go` - Gin router setup
- `cmd/` - Application entry point
- `test/` - Automated tests

## Learning Resources

This project is inspired by the [Go REST API Roadmap](https://roadmap.sh/golang/rest-api) tutorial, which provided a solid foundation for building RESTful APIs in Go. While I referenced the roadmap for best practices and overall guidance, I have customized and restructured the repository to fit my own learning goals and preferences. If you're starting out with Go or REST APIs, I highly recommend checking out the roadmap as a starting point!
