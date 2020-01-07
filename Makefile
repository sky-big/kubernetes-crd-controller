# Build all binary
#
# Example:
#   make
#   make all
all: build
.PHONY: all

# Build binary
#
# Example:
#   make build
build:
	go build -v -o ./bin/controller ./cmd/controller
.PHONY: build

# Generate code
#
# Example:
#   make generate
generate:
	hack/codegen.sh
.PHONY: generate

# Generate vendor
#
# Example:
#   make vendor
vendor:
	go mod vendor
.PHONY: vendor

# Clean package
#
# Example:
#   make clean
clean:
	rm -rf ./bin
