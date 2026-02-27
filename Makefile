# -------------------------------
# Makefile for Go project
# Automates formatting, linting, vetting, and tests
# -------------------------------

# Go binary
GO ?= go

# Linting tool
GOLANGCI_LINT ?= golangci-lint

# Directories
PKGS := $(shell $(GO) list ./...)

# Default target
.PHONY: all
all: fmt lint vet test

# -------------------------------
# Formatting
# -------------------------------
.PHONY: fmt
fmt:
	@echo "==> Running gofmt..."
	@gofmt -s -w $(PKGS)
	@echo "==> Running goimports..."
	@go install golang.org/x/tools/cmd/goimports@latest
	@goimports -w $(PKGS)

# -------------------------------
# Linting
# -------------------------------
.PHONY: lint
lint:
	@echo "==> Running golangci-lint..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@$(GOLANGCI_LINT) run --fix

# -------------------------------
# Go vet (static analysis)
# -------------------------------
.PHONY: vet
vet:
	@echo "==> Running go vet..."
	@$(GO) vet ./...

# -------------------------------
# Run tests
# -------------------------------
.PHONY: test
test:
	@echo "==> Running tests..."
	@$(GO) test -v ./...

# -------------------------------
# Clean build artifacts
# -------------------------------
.PHONY: clean
clean:
	@echo "==> Cleaning..."
	@rm -rf ./bin ./vendor

# -------------------------------
# Install dependencies
# -------------------------------
.PHONY: deps
deps:
	@echo "==> Installing dependencies..."
	@$(GO) mod tidy
	@$(GO) mod vendor

# -------------------------------
# Proto code generation (buf)
# -------------------------------
.PHONY: proto-gen
proto-gen:
	@echo "==> Generating proto code..."
	@buf dep update
	@buf generate
	@echo "==> Proto code generated."

.PHONY: proto-lint
proto-lint:
	@echo "==> Linting proto files..."
	@buf lint

.PHONY: proto-clean
proto-clean:
	@echo "==> Cleaning generated proto code..."
	@rm -f proto/*.pb.go proto/*_grpc.pb.go proto/*.pb.gw.go
	@echo "==> Proto code cleaned."

# -------------------------------
# Run all checks
# -------------------------------
.PHONY: check
check: fmt lint vet test
	@echo "==> All checks passed!"
