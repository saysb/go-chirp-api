# go-chirp

A minimalist Twitter clone API built with Go. Features text-only posts and comments. Learning project focused on Go backend development.

## Prerequisites

- Go 1.21+
- PostgreSQL 14+
- [golang-migrate/migrate](https://github.com/golang-migrate/migrate) for database migrations

## Configuration

Create a `.env` file with the following variables:

```env
# Security
PASSWORD_SECRET=your_secret_passphrase

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_postgres_username
DB_PASSWORD=your_postgres_password
DB_NAME=your_postgres_db_name
DB_SSLMODE=disable

# Server
PORT=8080
```

## Database Setup

The project uses PostgreSQL as its database. Make commands are provided for database management:

```bash
# Reset database
make db-reset

# Run migrations
make migrate-up

# Check migrations status
make migrate-status

# Create new migration
make migrate-create name=migration_name

# List database tables
make db-tables
```

## Running the Project

Start the server:

```bash
make run
```

The API will be available at `http://localhost:8080`.

## Project Structure

```
.
├── cmd/
│ └── api/ # Point d'entrée de l'application
├── internal/
│ ├── database/ # Couche d'accès aux données
│ ├── handlers/ # Gestionnaires HTTP
│ ├── models/ # Entity & DTO
│ ├── services/ # Logique metier avant insertion bdd
│ └── utils/ # Utilitaires
├── migrations/ # Fichiers de migration SQL
├── .env # Configuration
├── Makefile # Commandes make
└── README.md
```

## Available Make Commands

- `migrate-up`: Apply all pending migrations
- `migrate-down`: Rollback all migrations
- `migrate-create`: Create a new migration file
- `migrate-status`: Show current migration version
- `db-reset`: Drop and recreate database
- `db-tables`: List all tables in database
- `run`: Start the API server

## License

MIT