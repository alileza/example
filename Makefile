project_name = example
branch = $(shell git symbolic-ref HEAD 2>/dev/null)
version = $(shell cat Dockerfile | grep version | cut -d= -f2)
revision = $(shell git log -1 --pretty=format:"%H")
build_user = $(USER)
build_date = $(shell date +%FT%T%Z)
pwd = $(shell pwd)

module_name=$(shell head -n 1 go.mod | sed s/.......//)
version_pkg=$(module_name)/version
ldflags := "-X $(version_pkg).AppName=$(project_name) -X $(version_pkg).Version=$(version) -X $(version_pkg).Branch=$(branch) -X $(version_pkg).Revision=$(revision) -X $(version_pkg).BuildUser=$(build_user) -X $(version_pkg).BuildDate=$(build_date)"

.PHONY: all build run test test-unit test-cov

build:	
	@go build -mod=vendor -ldflags $(ldflags) -o ./bin/$(project_name) .

build-fe:
	@cd ui && yarn build

run:
	@./bin/$(project_name) serve

gen:
	@rm -rf ./autogen
	@docker-compose run --rm generate

test: 
	@echo ">> running unit test"
	@make test-unit

test-unit:
	@echo ">> running unit test"
	@go test -mod vendor -coverprofile unit-coverage.out -race -cover -covermode=atomic $(shell go list ./...)

test-cov:
	@go test $(module_name)/${PKG} -coverprofile=coverage.out && go tool cover -html=coverage.out && rm coverage.out

test-tomato:
	@docker-compose -f ./tomato/docker-compose.yml run --rm tomato
	@make test-tomato-cleanup

test-tomato-run:
	@docker-compose -f ./tomato/docker-compose.yml run tomato ; make test-tomato-cleanup

test-tomato-cleanup:
	@docker-compose -f ./tomato/docker-compose.yml down
	@docker-compose -f ./tomato/docker-compose.yml rm
