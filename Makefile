 .ONESHELL:
GOFILES := $(shell find . -name '*.go' -not -path "./vendor/*" | egrep -v "^\./\.go" | grep -v _test.go)
DATE = $(shell date +'%s')

.PHONY: build

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	OPEN = xdg-open
endif
ifeq ($(UNAME_S),Darwin)
	OPEN = open
endif

gobuild:
	@CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -o build/bin/signer-hashicorp-vault-plugin

lint-ci: ## Check linting
	@misspell -error $(GOFILES)
	@golangci-lint run

prod: gobuild
	@docker-compose -f docker-compose.yml up --build vault

dev: build-linux docker-build
	@docker-compose -f docker-compose.dev.yml up --build vault

down:
	@docker-compose -f docker-compose.dev.yml down --volumes --timeout 0

docker-build:
	@DOCKER_BUILDKIT=1 docker build -t signer-hashicorp-vault-plugin .

docker-remove:
	rm -f build/bin/signer-hashicorp-vault-plugin
	docker exec -it signer-hashicorp-vault-plugin-vault-1 rm /vault/plugins/signer-hashicorp-vault-plugin

build-linux:
	go mod tidy
	env GOOS=linux GOARCH=amd64 go build -v -o ./build/bin/signer-hashicorp-vault-plugin
