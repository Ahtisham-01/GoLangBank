build:
	go build -o bin/gobank main.go api.go types.go

run: build
	@./bin/gobank

test:
	@go test -v ./...
