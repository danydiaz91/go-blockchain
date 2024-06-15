build:
	@go build -o ./bin/blockchain ./cmd

run: build
	@ ./bin/blockchain

test:
	@ go test -v ./...