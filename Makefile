HELLO_KUBIC_BIN := bin/hello-kubic

GO ?= go
GO_MD2MAN ?= go-md2man

VERSION := $(shell cat VERSION)
USE_VENDOR =
LOCAL_LDFLAGS = -buildmode=pie -ldflags "-X=main.version=$(VERSION)"

.PHONY: all api build vendor
all: dep build

dep: ## Get the dependencies
	@$(GO) get -v -d ./...

update: ## Get and update the dependencies
	@$(GO) get -v -d -u ./...

tidy: ## Clean up dependencies
	@$(GO) mod tidy

vendor: dep ## Create vendor directory
	@$(GO) mod vendor

build: ## Build the binary files
	$(GO) build -v -o $(HELLO_KUBIC_BIN) $(USE_VENDOR) $(LOCAL_LDFLAGS) ./cmd/hello-kubic

clean: ## Remove previous builds
	@rm -f $(HELLO_KUBIC_BIN)

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


.PHONY: release
release: ## create release package from git
	git clone https://github.com/thkukuk/hello-kubic
	mv hello-kubic hello-kubic-$(VERSION)
	#sed -i -e 's|USE_VENDOR =|USE_VENDOR = -mod vendor|g' hello-kubic-$(VERSION)/Makefile
	#make -C hello-kubic-$(VERSION) vendor
	tar --exclude .git -cJf hello-kubic-$(VERSION).tar.xz hello-kubic-$(VERSION)
	rm -rf hello-kubic-$(VERSION)
