.DEFAULT_GOAL=help
GOOS=$(shell uname | tr '[:upper:]' '[:lower:]')
RELEASES_PATH=release
COLOR_NONE="\\033[0m"
COLOR_BLUE="\\033[34m"
COLOR_CYAN="\\033[36m"
COLOR_GREEN="\\033[32m"
COLOR_YELLOW="\\033[33m"
COLOR_ORANGE="\\033[43m"
COLOR_RED="\\033[31m"

help: #? help me
	@printf "\e[34;01mAvailable targets\033[0m\n"
	@awk '/^@?[a-zA-Z\-_0-9]+:/ { \
		nb = sub( /^#\? /, "", helpMsg ); \
		if(nb == 0) { \
			helpMsg = $$0; \
			nb = sub( /^[^:]*:.* #\? /, "", helpMsg ); \
		} \
		if (nb) \
			printf "\033[1;31m%-" width "s\033[0m %s\n", $$1, helpMsg; \
	} \
	{ helpMsg = $$0 }' \
	$(MAKEFILE_LIST) | column -ts:

build: #? install dependencies to start contributing
	go mod download

install: package #? package and install app to system
	mv $(RELEASES_PATH)/skeleton-$(GOOS) /usr/local/bin/skeleton

lint: lint-mod lint-cs #? pre-build validations
	@echo "$(COLOR_GREEN)All pre-build validations OK!$(COLOR_NONE)"
lint-mod:
	go mod verify
lint-cs:
ifneq ($(shell gofmt -e -l **/*.go),)
	@echo "$(COLOR_YELLOW)Files with code style errors:$(COLOR_NONE)"
	@echo $(shell gofmt -e -l **/*.go)
	@echo "$(COLOR_RED)Run \`make fix-cs\` to fix code style$(COLOR_NONE)"
endif

fix-cs: #? fix code style
	go fmt ./...

package: build #? create artifact for later usage
	go build -race -ldflags "-s -w -extldflags '-static'" -o $(RELEASES_PATH)/skeleton-$(GOOS)

test: #? test application
	go vet ./...
	go test -race -coverprofile=c.out ./...
