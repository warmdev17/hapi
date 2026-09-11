-include .env
export

run:
	go run cmd/server/main.go
sqlc:
	sqlc generate
migrate-up:
	goose -dir ./internal/db/migrations postgres "postgres://$$DB_USER:$$DB_PASSWORD@127.0.0.1:$$DB_PORT/$$DB_NAME?sslmode=disable" up
migrate-down:
	goose -dir ./internal/db/migrations postgres "postgres://$$DB_USER:$$DB_PASSWORD@127.0.0.1:$$DB_PORT/$$DB_NAME?sslmode=disable" down
