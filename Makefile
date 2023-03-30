SHELL=/bin/sh

export GO111MODULE=on

GOOS := $(or $(GOOS), '')
app := ${APP_NAME}

.PHONY: test
test:
	go test -cover -race -count=1 $(shell go list -f '{{.Dir}}/...' -m | xargs)

.PHONY: deps
deps:
	go mod download
	go mod tidy


.PHONY: app-deps
app-deps:
	cd ./${app}
	go mod download
	go mod tidy
	cd ..


.PHONY: lint
lint:
	golangci-lint run $(shell go list -f '{{.Dir}}/...' -m | xargs)

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=${GOOS} go build -o bin/${app} ${app}/cmd/*.go
