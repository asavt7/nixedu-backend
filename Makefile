include .env
export

.PHONY: build
build:
	go build -v -o ./bin/main ./cmd/main.go


.PHONY: docker-build
docker-build:
	docker build -t ${PROJECT_NAME} .


.PHONY: run
run:
	go run ./cmd/main.go


.PHONY: cover
cover:
	go test -race -v -coverprofile=./report/coverage.out -cover `go list ./... | grep -v mocks`
	go tool cover -func=./report/coverage.out
	go tool cover -html=./report/coverage.out


MOCKS_DESTINATION=mocks
.PHONY: mocks
mocks: $(shell find ./pkg -name "*.go")
	@echo "Generating mocks..."
	@rm -rf $(MOCKS_DESTINATION)
	@for file in $^; do mockgen -source=$$file -destination=$(MOCKS_DESTINATION)/$$file; done


.PHONY: test
test: mocks
	go test -v -race -timeout 30s ./...


.PHONY: migrate-up
migrate-up:
	migrate -path ./migrations/ -database 'postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}' up

## linter - run linter for go code
.PHONY: linter
linter:
	revive -config .linter.config.toml -formatter unix ./...

.PHONY: migrate-down
migrate-down:
	migrate -path ./migrations/ -database 'postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}' down


.PHONY: docker-compose-dev-up
docker-compose:
	docker-compose up db redis

.PHONY: swagger
swagger:
	swag init  --dir ./cmd,./pkg --parseInternal true

.DEFAULT_GOAL := build
