HOSTNAME=registry.terraform.io
NS=marcofranssen
NAME=curl
GIT_VERSION ?= $(shell git describe --tags --always --dirty)
# LDFLAGS="-X $(PKG).GitVersion=$(GIT_VERSION) -X $(PKG).gitCommit=$(GIT_HASH) -X $(PKG).gitTreeState=$(GIT_TREESTATE) -X $(PKG).buildDate=$(BUILD_DATE)"

# GO_BUILD_FLAGS := -trimpath -ldflags $(LDFLAGS)
GO_BUILD_FLAGS := -trimpath
COMMAND       := terraform-provider-$(NAME)

OS_ARCH=darwin_amd64
# INSTALL_DIR=~/.terraform.d/plugin-cache/$(HOSTNAME)/$(NS)/$(NAME)/$(GIT_VERSION)/$(OS_ARCH)
# INSTALL_DIR=~/.terraform.d/plugin-cache/$(HOSTNAME)/$(NS)/$(NAME)/0.1.0/$(OS_ARCH)
INSTALL_DIR=~/.terraform.d/plugins/$(HOSTNAME)/$(NS)/$(NAME)/0.1.0/$(OS_ARCH)

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
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}'

bin/terraform-provider-curl_%_$(GIT_VERSION): cmd/terraform-provider-curl
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 go build $(GO_BUILD_FLAGS) -o $@ ./$<

$(BINARIES): GOOS = $(word 1,$(subst _, ,$*))
$(BINARIES): GOARCH = $(word 2,$(subst _, ,$*))
build: $(BINARIES) ## Build the binary

clean-bin: ## Remove compiled binaries
	rm -r bin

$(INSTALL_DIR):
	@mkdir -p $@

$(INSTALL_DIR)/$(COMMAND): $(INSTALL_DIR)
	@echo Plugin installed at $@
	@cp bin/terraform-provider-curl_$(OS_ARCH)_$(GIT_VERSION) $@

install: build clean-install clean-plugin-cache $(INSTALL_DIR)/$(COMMAND) ## Install the plugin

clean-install: ## Removes the installed binary
	@echo Cleaning $(INSTALL_DIR)…
	@rm -r $(INSTALL_DIR) 2>/dev/null || true

clean-plugin-cache: ## Removes the installed binary
	@echo Cleaning $(subst plugins,plugin-cache,$(INSTALL_DIR))…
	@rm -r $(subst plugins,plugin-cache,$(INSTALL_DIR)) 2>/dev/null || true
