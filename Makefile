.ONESHELL:
GOFILES := $(shell find . -name '*.go' -not -path "./vendor/*" | egrep -v "^\./\.go" | grep -v _test.go)
DATE = $(shell date +'%s')

.PHONY: build lint

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	OPEN = xdg-open
endif
ifeq ($(UNAME_S),Darwin)
	OPEN = open
endif

test:
	@mkdir -p build/coverage
	go test  ./... -cover -coverprofile=build/coverage/coverage.out -covermode=count

run-coverage: test
	@sh scripts/coverage.sh build/coverage/coverage.out build/coverage/coverage.html

coverage: run-coverage
	@$(OPEN) build/coverage/coverage.html 2>/dev/null

gobuild:
	@CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -o build/bin/quorum-hashicorp-vault-plugin

lint-tools: ## Install linting tools
	@GO111MODULE=on go get github.com/client9/misspell/cmd/misspell@v0.3.4
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.27.0

lint:
	@misspell -w $(GOFILES)
	@golangci-lint run --fix

lint-ci: ## Check linting
	@misspell -error $(GOFILES)
	@golangci-lint run

prod: gobuild
	@docker-compose -f docker-compose.yml up --build vault

dev:  gobuild docker-build
	@docker-compose -f docker-compose.dev.yml up --build vault

down:
	@docker-compose -f docker-compose.dev.yml down --volumes --timeout 0

docker-build:
	@DOCKER_BUILDKIT=1 docker build -t quorum-hashicorp-vault-plugin .
