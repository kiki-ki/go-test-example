run:
	go run ./cmd/main.go

fmt:
	go vet ./...
	go fmt ./...

get-hello:
	curl localhost:8080/hello
