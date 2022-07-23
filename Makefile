migrateup:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/brodo_demo?sslmode=disable" -verbose up
migratedown:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/brodo_demo?sslmode=disable" -verbose down
.PHONY: migrateup migratedown