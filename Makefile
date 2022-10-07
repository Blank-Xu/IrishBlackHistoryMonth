# Usage:
# make
# make build
# make run
# make test

SERVICE_NAME=webservice

build:
	@echo "build start..."
	@echo " - download the necessary packages"
	go mod download
	@echo " - verify the packages"
	go mod verify
	@echo " - build the service"
	CGO_ENABLED=0 go build -o ${SERVICE_NAME}
	@echo "build success..."

run:
	@echo "run service..."
	go run main.go

test:
	@echo "test start..."
	go test -v  -cover ./...

.PHONY: build run test