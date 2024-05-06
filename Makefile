# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
BINARY_NAME = go-setup

# Targets
all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -rf ./bin

run:
	$(GOBUILD) -o ./bin/$(BINARY_NAME) -v 
	./bin/$(BINARY_NAME)

deps:
	$(GOGET) mod tidy

migration_up: 
	migrate -path database/migration/ -database "$(DATABASE_URL)" -verbose up

migration_down: 
	migrate -path database/migration/ -database "$(DATABASE_URL)" -verbose down

migration_fix: 
	migrate -path database/migration/ -database "$(DATABASE_URL)" force VERSION

lint:
	golangci-lint run

.PHONY: all build test clean run lint
