# Overview

This is a simple API that lets you do CRUD operations for User, and dynamically calculate their age using DOB. Age is not persisted in the database.

**Tech Stack**
- Fiber
- PostgreSQL
- sqlc
- Zap
- go migrate

**Note**: Dates are handled using a custom Date type with the format: YYYY-MM-DD.
All endpoints require  a `Content-Type` header, so make sure it is set

## API Endpoints

1. Create a User
```bash
POST /users
```

Request Body (Date format: YYYY-MM-DD):
```json
{
  "name": "John Doe",
  "dob": "2000-01-01" 
}
```

Response:
```json
{
  "id": 1,
  "name": "John Doe",
  "dob": "2000-01-01"
}
```

2. Get User by ID
```bash
GET /users/:id
```

Response:
```json
{
  "id": 1,
  "name": "John Doe",
  "dob": "2000-01-01",
  "age": 25
}
```

3. Update User
```bash
PUT /users/:id
```

Request Body:
```json
{
  "name": "Jane Doe",
  "dob": "1999-12-31"
}
```

Response:
```json
{
  "id": 1,
  "name": "Jane Doe",
  "dob": "1999-12-31"
}
```

4. Delete a User

```bash
DELETE /users/:id
```

Response: No Content

## How to Set up

1. Clone the repository
```bash
git clone git@github.com:Infamous003/ainyx.git
cd ainyx
```

2. Configure environment variables

- `DB_DSN` - The postgreSQL Data Source Name
example: `postgres://user:password@localhost:5432/dbname?sslmode=disable`
- `PORT`(Optional) - The server/API port(default is `8080`)

Example .env file:
```bash
export DB_DSN=postgres://postgres:postgres@localhost:5432/usersdb?sslmode=disable
export PORT=8080
```
Restart your terminal.

3. Install `go-migrate`
```bash
go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Apply `up` migrations:
```bash
migrate -database "$DB_DSN" -path db/migrations up
```

For down migrations(if needed):
```bash
migrate -database "$DB_DSN" -path db/migrations down 1
```

4. Install dependencies and run the server
```bash
go mod tidy
go run ./cmd/server
```
