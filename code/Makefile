ifdef VERSION
    VERSION=${VERSION}
else
    VERSION=$(shell git describe --always)
endif

ifdef GITCOMMIT
    GITCOMMIT=${GITCOMMIT}
else
    GITCOMMIT=$(shell git rev-parse HEAD)
endif

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

BUILDTIME=${shell date +%Y-%m-%dT%H:%M:%S%z}

GINKGO ?= $(LOCALBIN)/ginkgo
GOLINES ?= $(LOCALBIN)/golines
GOFUMPT ?= $(LOCALBIN)/gofumpt
GOCILINT ?= $(LOCALBIN)/golangci-lint

.PHONY: help
help: ## 展示可用 make 命令及说明
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


##@ 开发常用命令

.PHONY: build
build: tidy # build executable binary
	go build -o chatgpt-bot ./code/main.go

.PHONY: fmt
fmt: ## 执行 golines，gofumpt ...
	$(GOLINES) ./ -m 119 -w --base-formatter gofmt --no-reformat-tags
	$(GOFUMPT) -l -w .

.PHONY: vet
vet: ## 执行 go vet ./...
	go vet ./...

.PHONY: tidy
tidy: ## 执行 go mod tidy
	go mod tidy

.PHONY: test
test: tidy ## 执行 ginkgo 单元测试
	go test ./...

.PHONY: lint
lint: ## 执行 golangci-lint run
	$(GOCILINT) run

##@ 开发工具安装命令

.PHONY: install-golines
install-golines: $(GOLINES) ## Download golines locally if necessary
$(GOLINES): $(LOCALBIN)
	GOBIN=$(LOCALBIN) go install github.com/segmentio/golines@v0.11.0

.PHONY: install-gofumpt
install-gofumpt: $(GOFUMPT) ## Download gofumpt locally if necessary
$(GOFUMPT): $(LOCALBIN)
	GOBIN=$(LOCALBIN) go install mvdan.cc/gofumpt@v0.4.0

.PHONY: install-golangci-lint
install-golangci-lint: $(GOCILINT) ## Download golangci-lint locally if necessary
$(GOCILINT): $(LOCALBIN)
	GOBIN=$(LOCALBIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.47.3
