# Build binary
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
    go build -v -o ./bin/service ./cmd/service
.PHONY: build

# Clean package
#
# Example:
#   make clean
clean:
	rm -rf ./bin
