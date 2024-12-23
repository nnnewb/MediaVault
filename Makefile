GO?=go1.23.3
BINARY_DIR=bin
ARTIFACT=media-vault
ARTIFACT_SUFFIX=

GOOS?=$(shell $(GO) env GOOS)
GOARCH?=$(shell $(GO) env GOARCH)
CGO_ENABLED?=0
GOENV=CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH)

ifeq ($(GOOS),windows)
	ARTIFACT_SUFFIX=.exe
endif

.PHONY: build
build:
	$(GOENV) \
	$(GO) build \
		-trimpath \
		-o $(BINARY_DIR)/$(ARTIFACT)$(ARTIFACT_SUFFIX) \
		main.go
