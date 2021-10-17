.PHONY: build
build:
	go build -v -o ./bin/main ./cmd/main.go

MOCKS_DESTINATION=mocks
.PHONY: mocks
mocks: $(shell find ./pkg -name "*.go")
	@echo "Generating mocks..."
	@rm -rf $(MOCKS_DESTINATION)
	@for file in $^; do mockgen -source=$$file -destination=$(MOCKS_DESTINATION)/$$file; done

.DEFAULT_GOAL := build

.PHONY: test
test: mocks
	go test -v -race -timeout 30s ./...
