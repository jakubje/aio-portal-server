postgres:
	docker run --name portal -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres
createdb: 
	docker exec -it portal createdb --username=root --owner=root aio_portal

dropdb:
	docker exec -it portal dropdb aio_portal

migrateup:
	migrate -path ./db/migrations -database "postgresql://root:root@localhost:5432/aio_portal?sslmode=disable" -verbose up

migratedown:
	migrate -path ./db/migrations -database "postgresql://root:root@localhost:5432/aio_portal?sslmode=disable" -verbose down
 
sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server
 
