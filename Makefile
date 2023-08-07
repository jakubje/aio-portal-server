postgres:
	docker run --name portal --network aioportal-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres

createdb: 
	docker exec -it portal createdb --username=root --owner=root aio_portal

dropdb:
	docker exec -it portal dropdb aio_portal

migrateup:
	migrate -path ./db/migrations -database "postgresql://root:root@localhost:5432/aio_portal?sslmode=disable" -verbose up

migrateup1:
	migrate -path ./db/migrations -database "postgresql://root:root@localhost:5432/aio_portal?sslmode=disable" -verbose up 1

migratedown:
	migrate -path ./db/migrations -database "postgresql://root:root@localhost:5432/aio_portal?sslmode=disable" -verbose down

migratedown1:
	migrate -path ./db/migrations -database "postgresql://root:root@localhost:5432/aio_portal?sslmode=disable" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migrations -seq $(name)

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/jakub/aioportal/server/db/sqlc Store

proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
        --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
        --openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=aio_portal \
        proto/*.proto
	statik -src=./doc/swagger -dest=./doc

evans:
	evans --host localhost --port 9090 -r repl

redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 new_migration db_schema sqlc test server mock proto evans redis

 
