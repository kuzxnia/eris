
runserver:
	air --build.cmd "go build -o bin/api main.go" --build.bin "./bin/api"

test:
	gotestsum ./...

lint:
	golangci-lint run

swagger:
	swag fmt
	swag init

gen: 
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/interfaces/web/proto/*.proto

build-dev:
	# only your platform
	goreleaser build --snapshot --clean --single-target -f agent/.goreleaser.yaml

build-all:
	# only your platform
	goreleaser build --snapshot --clean

.PHONY: runserver lint swagger test gen build-dev build-all build-chart
