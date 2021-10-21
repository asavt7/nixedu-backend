include .env
export

PROJECT_NAME=nix-edu-backend

## ----------------------------------------------------------------------
## 		A little manual for using this Makefile.
## ----------------------------------------------------------------------


.PHONY: build
build:	## Compile the code into an executable application
	go build -v -o ./bin/main ./cmd/main.go


.PHONY: docker-build
docker-build:	## Build docker image
	go mod vendor
	docker build -t ${PROJECT_NAME} .


.PHONY: run
run:	## Run application
	go run ./cmd/main.go


.PHONY: cover
cover:	## run tests and show test coverage
	go test -race -v -coverprofile=./report/coverage.out -cover `go list ./... | grep -v mocks`
	go tool cover -func=./report/coverage.out
	go tool cover -html=./report/coverage.out


MOCKS_DESTINATION=mocks
.PHONY: mocks
mocks: ## Generate mocks
	@echo "Generating mocks..."
	@rm -rf $(MOCKS_DESTINATION)
	@for file in $(shell find ./pkg -name "*.go" | xargs echo); do mockgen -source=$$file -destination=$(MOCKS_DESTINATION)/$$file; done


.PHONY: test
test: mocks ## Run golang tests
	go test -v -race -timeout 30s ./...


.PHONY: migrate-up
migrate-up:	## run db migration scripts
	migrate -path ./migrations/ -database 'postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}' up

.PHONY: migrate-down
migrate-down:	## rollback db migrations
	migrate -path ./migrations/ -database 'postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}' down

.PHONY: linter
linter:	## Run linter for *.go files
	revive -config .linter.config.toml -formatter unix ./...

.PHONY: docker-compose-dev-up
docker-compose-dev-up:	## Run local dev environment
	docker-compose up db redis

.PHONY: swagger
swagger:	## Generate swagger api specs
	swag init  --dir ./cmd,./pkg --parseInternal true


.PHONY: docs
docs:	## Run godoc
	@echo  'link to docs  http://localhost:6060/pkg/github.com/asavt7/nixEducation/'
	godoc -http=:6060


.PHONY: help
help:     ## Show this help.
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)


.DEFAULT_GOAL := build
