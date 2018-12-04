# Tic Tac Go
Tic Tac Toe game implemented in Go.

## Getting Started
Install go 1.11.2 (e.g. using Homebrew or gvm).

## Project Structure
This project follows the [recommended Go project structure](https://github.com/golang-standards/project-layout). In particular:

- `cmd` contains the tictacgo main executable
- `internal` contains code for the implementation of tictacgo and accompanying unit tests
- `test` contains "external" test code (like acceptance tests)

This project is also setup as the `github.com/bgerstle/tictacgo` module. As a result, running commands which target subdirectories of this project must contain the module path. For example, when building:

## Building The App
To compile a static executable, use `go build`:

``` shell
go build github.com/bgerstle/tictacgo/cmd/tictacgo
```

The result will be in the root directory, which you can run via: `./tictacgo`.

## Running The App
Build and run the app as described above, or you can run from source using `go run`:

``` shell
go run github.com/bgerstle/tictacgo/cmd/tictacgo/main.go
```

## Running The Tests
Tests are run using `go test`, using the "recursive import path" specifier (`./...`):

- All tests: `go test ./...`
- Unit tests: `go test ./internal/...`
- Acceptance Tests: `go test ./test/acceptance/...`

## Linting The Code
Code is linted using `golint`

``` shell
golint ./...
```