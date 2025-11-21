# ===========================
# Operation GoGoGo — Makefile
# ===========================

APPLICATION_NAME = operation-gogogo
MAIN_PACKAGE = ./cmd/api

# Version: git commit hash (short)
VERSION = $(shell git describe --tags --always)

# Build the application with the version injected
build:
	@echo "Building $(APPLICATION_NAME) (version: $(VERSION))..."
	@go build \
		-ldflags "-X main.buildVersion=$(VERSION)" \
		-o bin/$(APPLICATION_NAME) \
		$(MAIN_PACKAGE)
	@echo "Build completed → bin/$(APPLICATION_NAME)"

# Run the application directly (with auto version)
run:
	@echo "Running $(APPLICATION_NAME) (version: $(VERSION))..."
	@go run \
		-ldflags "-X main.buildVersion=$(VERSION)" \
		$(MAIN_PACKAGE)

# Run tests
test:
	@echo "Running tests..."
	@go test ./...

# Remove compiled binary
clean:
	@echo "Cleaning build output..."
	@rm -rf bin
	@echo "Done."

# Format code
fmt:
	@go fmt ./...

# Tidy modules
tidy:
	@go mod tidy
