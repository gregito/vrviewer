GOFILES_NOVENDOR = $(shell find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./.git/*")
BINARY=vrviewer

build: build-darwin build-linux build-windows

format:
	@gofmt -w ${GOFILES_NOVENDOR}

build-darwin:
	GOOS=darwin GO111MODULE=on CGO_ENABLED=0 go build -mod=vendor -o build/Darwin/${BINARY} main.go

build-linux:
	GOOS=linux GO111MODULE=on CGO_ENABLED=0 go build -mod=vendor -o build/Linux/${BINARY} main.go

build-windows:
	GOOS=windows GO111MODULE=on CGO_ENABLED=0 go build -mod=vendor -o build/Windows/${BINARY}.exe main.go

test:
	GO111MODULE=on go test -mod=vendor -timeout 30s -race ./...

build-darwin-version:
	GOOS=darwin GO111MODULE=on CGO_ENABLED=0 go build -o build/Darwin/${BINARY} main.go

build-linux-version:
	GOOS=linux GO111MODULE=on CGO_ENABLED=0 go build -o build/Linux/${BINARY} main.go

build-windows-version:
	GOOS=windows GO111MODULE=on CGO_ENABLED=0 go build -o build/Windows/${BINARY}.exe main.go

install: build ## Installs OS specific binary into: /usr/local/bin
	install build/$(shell uname -s)/$(BINARY) /usr/local/bin

.DEFAULT_GOAL := build