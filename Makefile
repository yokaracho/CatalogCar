createDB:
	migrate -database postgres://postgres:12345@localhost:5432/postgres?sslmode=disable -path db/migrations up
dropDB:
	migrate -database postgres://postgres:12345@localhost:5432/postgres?sslmode=disable -path db/migrations drop