# GraphQL Server

A modern GraphQL server built with Go using the gqlgen framework.

## Features

- GraphQL API with playground interface
- Query caching with LRU cache
- Automatic persisted queries
- Docker support
- Built with gqlgen for type-safe GraphQL implementation

## Prerequisites

- Go 1.24.2 or later
- Docker (optional, for containerized deployment)

## Getting Started

### Local Development

1. Clone the repository:
   ```bash
   git clone https://github.com/sdshah09/Go-GraphQL.git
   cd Go-GraphQL
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run the server:
   ```bash
   go run server.go
   ```

4. Access the GraphQL playground:
   Open your browser and navigate to `http://localhost:8080`

### Docker Deployment

1. Build the Docker image:
   ```bash
   docker build -t my-graphql-server .
   ```

2. Run the container:
   ```bash
   docker run -p 8080:8080 my-graphql-server
   ```

## API Documentation

The GraphQL playground is available at `http://localhost:8080` when the server is running. This interactive interface allows you to:

- Explore the GraphQL schema
- Test queries and mutations
- View API documentation

## Project Structure

```
.
├── Dockerfile          # Docker configuration
├── graph/             # GraphQL schema and resolvers
├── gqlgen.yml         # gqlgen configuration
├── go.mod             # Go module dependencies
├── go.sum             # Go module checksums
├── server.go          # Main server implementation
└── tools.go           # Go tools configuration
```

## Dependencies

- [gqlgen](https://github.com/99designs/gqlgen) - Go GraphQL implementation
- [gqlparser](https://github.com/vektah/gqlparser) - GraphQL parser

