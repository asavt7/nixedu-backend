include .env
export

PROJECT_NAME=nixedu-backend

## ----------------------------------------------------------------------
## 		A little manual for using this Makefile.
## ----------------------------------------------------------------------


.PHONY: build
build:	## Compile the code into an executable application
	go build -v -o ./bin/main ./cmd/server/main.go


.PHONY: docker-build
docker-build:	## Build docker image
	docker build -t ${PROJECT_NAME} .


.PHONY: run
run:	## Run application
	go run ./cmd/server/main.go


.PHONY: cover
cover: test	## run tests and show test coverage
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out


MOCKS_DESTINATION=mocks
.PHONY: mocks
mocks: ## Generate mocks
	@echo "Generating mocks..."
	@rm -rf $(MOCKS_DESTINATION)
	@for file in $(shell find ./pkg -name "*.go" | xargs echo); do mockgen -source=$$file -destination=$(MOCKS_DESTINATION)/$$file; done


.PHONY: test
test: mocks ## Run golang tests
	go test -race  -coverprofile=coverage.out -cover `go list ./... | grep -v mocks `


.PHONY: migrate-up
migrate-up:	## run db migration scripts
	migrate -path ./migrations/ -database 'postgres://${PG_USERNAME}:${PG_PASSWORD}@${PG_HOST}:${PG_PORT}/${PG_NAME}?sslmode=${PG_SSLMODE}' up

.PHONY: migrate-down
migrate-down:	## rollback db migrations
	migrate -path ./migrations/ -database 'postgres://${PG_USERNAME}:${PG_PASSWORD}@${PG_HOST}:${PG_PORT}/${PG_NAME}?sslmode=${PG_SSLMODE}' down

.PHONY: linter
linter:	## Run linter for *.go files
	revive -config .linter.config.toml  -exclude ./vendor/... -formatter unix ./...


.PHONY: docker-compose-up
docker-compose-up:	## Run application and app environment in docker
	docker-compose up db redis backend-app migrate


.PHONY: docker-compose-dev-up
docker-compose-dev-up:	## Run local dev environment
	docker-compose up db redis migrate

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
