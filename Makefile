run:
	go run cmd/main.go

test:
	go test ./...

swagger:
	swag init --parseDependency --parseInternal -g cmd/main.go