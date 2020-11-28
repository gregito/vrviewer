GOFILES_NOVENDOR = $(shell find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./.git/*")

format:
	@gofmt -w ${GOFILES_NOVENDOR}