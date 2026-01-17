BINARY_NAME=hector
MAIN_FILE=main.go

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

.PHONY: all build test clean run deps tidy lint vet get

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN_FILE)

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run: build
	./$(BINARY_NAME)

deps:
	$(GOMOD) download

tidy:
	$(GOMOD) tidy

get:
	$(GOGET) ./...

lint:
	golangci-lint run

vet:
	$(GOCMD) vet ./...
