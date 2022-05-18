postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

dropdb:
	docker exec -it postgres12 dropdb simple_bank

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

wsqlc:
# following these step to run on window
# docker pull kjconroy/sqlc
# run this docker run --rm -v <absolute path to currrent directory>:/src -w /src kjconroy/sqlc generate
	docker run --rm -v G:\Studying\Go\SimpleBank\simplebank:/src -w /src kjconroy/sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc wsqlc test