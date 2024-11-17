DB_URL=postgresql://root:secret@localhost:5432/cars69?sslmode=disable

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine
createmigrateinitschema:
	 migrate create -ext sql --dir db/migrations --seq init_schema  

new_migration:
	migrate create -ext sql -dir db/migrations -seq $(name)

createdb:
	docker exec -it postgres createdb --username=root --owner=root cars69

dropdb:
	docker exec -it postgres dropdb cars69
sqlc:
	sqlc generate
server:
	go run main.go

test:
	go test -v -cover ./db/sqlc ./api ./token  
testmail:
	go test -v -cover ./email

migrateup:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migrations -database "$(DB_URL)" -verbose down

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/SabariGanesh-K/cars69-service.git/db/sqlc Store


redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

launch:
	docker start postgres
	docker start redis
.PHONY: postgres new_migration createdb dropdb migrateup createmigrateinitschema sqlc mock test  redis testmail launch