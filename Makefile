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


# ======================
# Swagger documentation
# ======================
# Requires: go install github.com/swaggo/swag/cmd/swag@latest
.PHONY: swagger swagger-init
swagger:
	@echo "Generating Swagger docs (if swag is installed)..."
	@command -v swag >/dev/null 2>&1 && swag init -g ./cmd/api/main.go -o ./docs || echo "swag not installed; skipping generation. Install with: go install github.com/swaggo/swag/cmd/swag@latest"

swagger-init: swagger
	@echo "Swagger docs are located at /swagger/index.html when the server is running."


# ========================================
# Version bumping (Semantic Versioning)
# ========================================

# Get the latest tag (or default to v0.0.0)
LATEST_TAG = $(shell git describe --tags --abbrev=0 2>/dev/null || echo v0.0.0)

# Extract MAJOR, MINOR, PATCH numbers
MAJOR = $(word 1,$(subst ., ,$(subst v,,$(LATEST_TAG))))
MINOR = $(word 2,$(subst ., ,$(subst v,,$(LATEST_TAG))))
PATCH = $(word 3,$(subst ., ,$(subst v,,$(LATEST_TAG))))

# ------------------------
# Bump PATCH
# ------------------------
tag-patch:
	@echo "Latest tag: $(LATEST_TAG)"
	@NEW_PATCH=$$(( $(PATCH) + 1 )); \
	NEW_VERSION=v$(MAJOR).$(MINOR).$$NEW_PATCH; \
	echo "Creating patch version: $$NEW_VERSION"; \
	git tag $$NEW_VERSION; \
	git push origin $$NEW_VERSION

# ------------------------
# Bump MINOR
# ------------------------
tag-minor:
	@echo "Latest tag: $(LATEST_TAG)"
	@NEW_MINOR=$$(( $(MINOR) + 1 )); \
	NEW_VERSION=v$(MAJOR).$$NEW_MINOR.0; \
	echo "Creating minor version: $$NEW_VERSION"; \
	git tag $$NEW_VERSION; \
	git push origin $$NEW_VERSION

# ------------------------
# Bump MAJOR
# ------------------------
tag-major:
	@echo "Latest tag: $(LATEST_TAG)"
	@NEW_MAJOR=$$(( $(MAJOR) + 1 )); \
	NEW_VERSION=v$$NEW_MAJOR.0.0; \
	echo "Creating major version: $$NEW_VERSION"; \
	git tag $$NEW_VERSION; \
	git push origin $$NEW_VERSION
