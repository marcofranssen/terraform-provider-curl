HOSTNAME=registry.terraform.io
NS=marcofranssen
NAME=curl
GIT_VERSION ?= $(shell git describe --tags --always --dirty)
# LDFLAGS="-X $(PKG).GitVersion=$(GIT_VERSION) -X $(PKG).gitCommit=$(GIT_HASH) -X $(PKG).gitTreeState=$(GIT_TREESTATE) -X $(PKG).buildDate=$(BUILD_DATE)"

# GO_BUILD_FLAGS := -trimpath -ldflags $(LDFLAGS)
GO_BUILD_FLAGS := -trimpath
COMMAND       := terraform-provider-$(NAME)

OS_ARCH=darwin_amd64

BUILDS=$(OS_ARCH) \
  darwin_arm64  \
  linux_386     \
  linux_amd64   \
  linux_arm     \
  linux_arm64   \
  windows_386   \
  windows_amd64

BINARIES=$(BUILDS:%=bin/$(COMMAND)_%_$(GIT_VERSION))

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Documentation:

$(GO_PATH)/bin/tfplugindocs:
	go install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@latest

.PHONY: provider-docs
provider-docs: $(GO_PATH)/bin/tfplugindocs ## Generate provider documentation
	@go generate ./...

$(GO_PATH)/bin/terraform-docs:
	go install github.com/terraform-docs/terraform-docs@latest

.PHONY: example-docs
example-docs: $(GO_PATH)/bin/terraform-docs ## Generates terraform documentation for examples
	@$< markdown examples/default --output-file README.md
	@$< markdown examples/ifconfig --output-file README.md
	@$< markdown examples/trigger-github-workflow --output-file README.md

##@ Testing:

.PHONY: test
test: ## Run unit tests
	go test -v -count=1 ./...

.PHONY: acc-test
acc-test: ## Run acceptance tests
	TF_ACC=1 go test -v -count=1 ./pkg/curl

##@ Install:

GOBIN := $(shell go env GOBIN)
ifeq ($(strip $(GOBIN)),)
    # If GOBIN is empty, set the variable GOBIN to the $GOPATH/bin bin folder
    GOBIN := $(shell go env GOPATH)/bin
endif

.PHONY: install
install: ## Install the provider to $GOBIN
	@echo Installing provider to $(GOBIN)…
	@go install .

bin/terraform-provider-curl_%_$(GIT_VERSION): .
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 go build $(GO_BUILD_FLAGS) -o $@ $<

##@ Build:

$(BINARIES): GOOS = $(word 1,$(subst _, ,$*))
$(BINARIES): GOARCH = $(word 2,$(subst _, ,$*))
build: $(BINARIES) ## Build the binary

clean-bin: ## Remove compiled binaries
	rm -r bin
