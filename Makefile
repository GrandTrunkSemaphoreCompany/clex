GOFILES = $(shell find . -name '*.go')
GOTESTFILES = $(shell find . -name '*.go')

default: build

workdir:
	mkdir -p bin

build: build-linux

build-linux: $(GOFILES)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/clex .

build-rasp: $(GOFILES)
	GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 go build -o bin/clex .

test: $(GOTESTFILES)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go test .
