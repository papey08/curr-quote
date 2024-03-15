test:
	go mod download && go test ./... ./...

run:
	docker-compose up
