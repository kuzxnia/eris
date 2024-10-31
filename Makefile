
runserver:
	air --build.cmd "go build -o bin/api main.go" --build.bin "./bin/api"

lint:
	golangci-lint run

swagger:
	swag fmt
	swag init

.PHONY: runserver lint swagger
