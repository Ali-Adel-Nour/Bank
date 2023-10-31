postgres:

	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine


createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

drobdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

makeFileDir := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

sqlc:
	docker run --rm -v $(makeFileDir):/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb drobdb migrateup migratedown sqlc test server
