# Go Empty Project

A minimal, extensible Go backend project using [Gin](https://github.com/gin-gonic/gin) and [PostgreSQL](https://www.postgresql.org/), designed as a clean starting point for new web services. It includes basic user management, JWT authentication scaffolding, Docker support, and Swagger API documentation.

---

## Features

- RESTful API with [Gin](https://github.com/gin-gonic/gin)
- PostgreSQL integration (via [pgx](https://github.com/jackc/pgx))
- User CRUD (Create, Read by ID)
- JWT authentication (scaffolded)
- CORS enabled
- Graceful shutdown
- Configurable via environment variables
- Docker & Docker Compose support
- Swagger/OpenAPI documentation ([Swaggo](https://github.com/swaggo/swag))

---

## Getting Started

### Prerequisites
- Go 1.23+
- PostgreSQL
- (Optional) Docker & Docker Compose

### Clone the repository
```sh
git clone <your-repo-url>
cd empty-go-project
```

### Install dependencies
```sh
go mod tidy
```

### Environment Variables
Create a `.env` file in the project root with the following variables:

```env
PORT=":8080"
DB_HOST=localhost
DB_PORT=5432
DB_USER=youruser
DB_PASSWORD=yourpassword
DB_NAME=yourdb
GIN_MODE=debug
LOGGER_FOLDER_PATH=./logs
LOGGER_FILENAME=app.log
ACCESS_KEY=your-access-secret
ACCESS_TIME=15m
REFRESH_KEY=your-refresh-secret
REFRESH_TIME=168h
APP_VERSION=1.0.0
UPLOAD_PATH=./uploads
```

### Database Setup
Create a `users` table in your PostgreSQL database:
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);
```

### Run the application
```sh
cd cmd/http
go run main.go
```

Or use the Makefile:
```sh
make dev
```

---

## API Endpoints

### User
- `POST /api/v1/users/` — Create a new user
    - Request body:
      ```json
      {
        "name": "John Doe",
        "email": "john@example.com",
        "password": "yourpassword"
      }
      ```
    - Response: `{ "message": "successfully created", "id": 1 }`

- `GET /api/v1/users/:id` — Get user by ID
    - Response:
      ```json
      {
        "id": 1,
        "name": "John Doe",
        "email": "john@example.com",
        "password": "<hashed>"
      }
      ```

### Swagger Docs
- [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

---

## Docker Usage

### Build and Run with Docker
```sh
docker build -t empty-go-app .
docker run --env-file .env -p 8080:8080 empty-go-app
```

### Docker Compose
Edit `docker-compose.yml` as needed, then:
```sh
docker-compose up --build
```

---

## Development

### Generate Swagger Docs
```sh
go install github.com/swaggo/swag/cmd/swag@latest
swag init -g ./cmd/http/main.go
```

### Lint, Test, Build
- Lint: `golint ./...`
- Test: `go test ./...`
- Build: `make build`

---

## Project Structure

```
cmd/http/           # Main entrypoint
internal/           # Application code (MVC structure)
  config/           # Configuration and logger
  delivery/http/    # HTTP handlers
  model/            # Data models
  repository/       # Data access layer
  route/            # Route definitions
  service/          # Business logic
  storage/postgres/ # DB connection and migrations
  utils/            # Utilities
pkg/                # Reusable packages (CORS, JWT, etc.)
docs/               # Swagger/OpenAPI docs
```

---

## Contributing

1. Fork the repo
2. Create your feature branch (`git checkout -b feature/foo`)
3. Commit your changes (`git commit -am 'Add feature'`)
4. Push to the branch (`git push origin feature/foo`)
5. Create a new Pull Request

---

## License

This project is licensed under the MIT License. 

<!-- Key Points of the MIT License
Freedom to Use: Anyone can use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the software.
Attribution Required: If you use the code, you must include the original copyright notice and the license text in any copies or substantial portions of the software.
No Warranty: The software is provided “as is”, without warranty of any kind. The authors are not liable for any damages or issues that arise from using the code.
Permissive: You can use the code in both open-source and proprietary (closed-source) projects. -->