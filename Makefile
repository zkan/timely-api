POSTGRESQL_URL=postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable

run:
	go run main.go

test:
	go test ./...

migrate-up:
	migrate -database ${POSTGRESQL_URL} -path db/migrations up

migrate-down:
	migrate -database ${POSTGRESQL_URL} -path db/migrations down
