.PHONY: build test clean run

# The name of your app
APPNAME := random-http-cats

# The path to the Go module
MODULE := github.com/fdpeiter/random-http-cats

# The Go workspace
GOPATH := $(shell go env GOPATH)

# The command to build the application
build:
	go build -o dist/$(APPNAME) ./app

# The command to run tests
test:
	go test ./...

# The command to clean up
clean:
	rm -f $(APPNAME)

run:
	go run ./app/main.go
