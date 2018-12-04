DEFAULT_GOAL: build

TTG_MODULE_NAME="github.com/bgerstle/tictacgo"

build:
	go build "$(TTG_MODULE_NAME)/cmd/tictacgo"
.PHONY: build

run:
	go run "$(TTG_MODULE_NAME)/cmd/tictacgo"
.PHONY: run

test: unit-test acceptance-test
.PHONY: test

unit-test:
	go test "$(TTG_MODULE_NAME)/internal/app/tictacgo"
.PHONY: unit-test

acceptance-test:
	go test "$(TTG_MODULE_NAME)/test/acceptance"
.PHONY: acceptance-test

lint:
	golint ./...
.PHONY: lint