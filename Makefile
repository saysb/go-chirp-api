# Makefile

# Variables
DB_NAME=twitter-clone
DB_USER=sebastiendamy
DB_PASSWORD=postgres
DB_HOST=localhost
DB_PORT=5432
DB_URL=postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

# Migrations
migrate-up:
	migrate -path migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" down

migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

migrate-status:
	@echo "Migrations status:"
	migrate -path migrations -database "$(DB_URL)" version

db-reset:
	dropdb $(DB_NAME) || true
	createdb $(DB_NAME)

db-tables:
	@echo "\nListe des tables dans la base de données:"
	@psql -d $(DB_NAME) -c "\dt"

run:
	go run cmd/api/main.go

.PHONY: migrate-up migrate-down migrate-create migrate-status db-reset db-tables