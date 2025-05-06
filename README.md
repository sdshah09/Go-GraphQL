# GraphQL Server with PostgreSQL and MongoDB

A GraphQL server that supports both PostgreSQL and MongoDB databases, built with Go and gqlgen.

## Features

- Dual database support (PostgreSQL and MongoDB)
- GraphQL API with playground
- CRUD operations for Patients
- Database schema validation
- Environment-based configuration

## Prerequisites

- Go 1.16 or higher
- PostgreSQL
- MongoDB
- Make (optional, for using Makefile commands)

## Setup

1. Clone the repository:
```bash
git clone https://github.com/sdshah09/Go-GraphQL
cd my-graphql-server
```

2. Install dependencies:
```bash
go mod download
```

3. Create a `.env` file in the root directory:
```env
# PostgreSQL Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=your_db_name

# MongoDB Configuration
MONGO_URI=mongodb://localhost:27017
MONGO_DB=your_db_name
```

4. Run the server:
```bash
go run server.go
```

The GraphQL playground will be available at `http://localhost:8080`

## API Examples

### Create a Patient

```graphql
mutation {
  createPatient(
    input: {
      first_name: "John"
      last_name: "Doe"
      db: "postgres"  # or "mongo"
    }
  ) {
    id
    first_name
    last_name
  }
}
```

### Query Patients

```graphql
query {
  patients(db: "postgres") {  # or "mongo"
    id
    first_name
    last_name
  }
}
```

## Database Schemas

### PostgreSQL
```sql
CREATE TABLE patients (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255)
);
```

### MongoDB
```javascript
{
  $jsonSchema: {
    bsonType: "object",
    required: ["first_name"],
    properties: {
      first_name: {
        bsonType: "string",
        description: "must be a string and is required"
      },
      last_name: {
        bsonType: "string",
        description: "must be a string if present"
      }
    }
  }
}
```

## Project Structure

```
.
├── config/         # Configuration management
├── graph/          # GraphQL schema and resolvers
├── utils/          # Database utilities
│   ├── postgres/   # PostgreSQL specific code
│   └── mongodb/    # MongoDB specific code
├── server.go       # Main server file
└── .env           # Environment variables
```

## Development

To regenerate GraphQL code after schema changes:
```bash
go run github.com/99designs/gqlgen generate
```
