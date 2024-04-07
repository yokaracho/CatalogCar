createDB:
	migrate -database postgres://postgres:password@localhost:5432/postgres?sslmode=disable -path db/migrations up
dropDB:
	migrate -database postgres://postgres:password@localhost:5432/postgres?sslmode=disable -path db/migrations drop
