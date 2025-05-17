run:
	go run cmd/main.go

test:
	go test ./...

swagger:
	swag init --parseDependency --parseInternal -g cmd/main.go

build-image:
	docker build . -t task-tracker:latest

run-system:
	docker-compose -p task-tracker-app up -d