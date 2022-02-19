help::
	@echo "======Available commands==========="
	@echo "make build"
	@echo "make test"

test:
	@echo "====================="
	@echo "Running unit tests"
	@echo "====================="
	go test -race ./...

build:
	@echo "====================="
	@echo "Building Project"
	@echo "====================="
	go build -o httptest
	@echo "Build complete"

