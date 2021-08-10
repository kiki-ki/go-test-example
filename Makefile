run:
	go run ./cmd/main.go

fmt:
	go vet ./...
	go fmt ./...
	golangci-lint run

db-seed:
	go run ./internal/interface/database/seed/seed.go

test-all:
	go test -v ./...

test-model:
	go test -v ./internal/app/model/...

test-handler:
	go test -v ./internal/app/handler/...

test-repository:
	go test -v ./internal/app/repository/...

hello:
	curl localhost:8080/hello

user-index:
	curl localhost:8080/users

user-show:
	curl localhost:8080/users/1

user-update:
	curl localhost:8080/users/1 -X PUT -H "Content-Type: application/json" -d '{"name":"chan", "email":"chan@exa.com", "age":100}'

user-create:
	curl localhost:8080/users -X POST -H "Content-Type: application/json" -d '{"name":"chan2", "email":"chan@exa.com", "age":200}'

user-delete:
	curl localhost:8080/users/1 -X DELETE

user-transaction:
	curl localhost:8080/users/transaction
