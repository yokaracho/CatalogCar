createDB:
	migrate -database postgres://postgres:11otozim@localhost:5432/postgres?sslmode=disable -path db/migrations up
dropDB:
	migrate -database postgres://postgres:11otozim@localhost:5432/postgres?sslmode=disable -path db/migrations drop