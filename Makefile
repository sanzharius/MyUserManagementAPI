include ./configs/.env
build:
	go build ./cmd/usermanager/main.go

run:
	go run ./cmd/usermanager/main.go

run-linter:
	echo "starting linters"
	golangci-lint run ./...




dev:
	echo "starting docker dev environment"
	docker-compose --env-file ./configs/.env -f deployments/docker-compose.yml up usermanager --build


