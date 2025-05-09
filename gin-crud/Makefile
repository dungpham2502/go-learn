postgres:
	docker run --name r -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres25 createdb --username=root --owner=root library_app

dropdb:
	docker exec -it postgres25 dropdb library_app

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/library_app?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/library_app?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc