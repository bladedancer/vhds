PROJECT_NAME := vhds
PKG_LIST := $(shell go list ./... | grep -v /vendor/)
REGISTRY ?= bladedancer
BIN ?= vhds

PKG := bladedancer/$(PROJECT_NAME)

.PHONY: clean helm

_all: clean lint helm docker-build push helm ## Build everything

all: vhds

vhds:
	@$(MAKE) _all BIN=vhds

lint: ## Lint the files
	@golint	-set_exit_status	${PKG_LIST}

build: ## Build the binary for linux
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build	 -o ./bin/$(BIN)	./cmd/$(BIN)

docker-build: ## Build docker image
	docker build -f ./Dockerfile -t $(REGISTRY)/$(BIN):latest	.

push: ## Push docker image
	docker push $(REGISTRY)/$(BIN):latest

helm: ## Update helm charts
	find ./helm -name *.tgz -exec rm {} \; && \
	helm dep update ./helm/envoyss && \
	helm dep update ./helm/back && \
	helm dep update ./helm/front && \
	helm dep update ./helm/xds && \
	helm dep update ./helm/ingress

clean: ## Clean out dir
	rm -rf ./bin && \
    docker rmi -f $(REGISTRY)/$(BIN):latest

help: ## Display this help screen
	@grep	-h	-E	'^[a-zA-Z_-]+:.*?## .*$$'	$(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
