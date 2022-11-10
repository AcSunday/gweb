CMD = go build
NAME = $(shell go list -m)
MAIN_GO = ./application
CGO_ENABLED = 0
BUILDTAGS = netgo
BUILDFLAGS = '-extldflags "-lm -lstdc++ -static"'

all: build

build:
	CGO_ENABLED=$(CGO_ENABLED) GOARCH=amd64 $(CMD) -tags $(BUILDTAGS) -ldflags $(BUILDFLAGS) -o build/$(NAME) $(MAIN_GO)

build-linux:
	CGO_ENABLED=$(CGO_ENABLED) GOOS=linux GOARCH=amd64 $(GO) build -tags $(BUILDTAGS) -ldflags $(BUILDFLAGS) -o build/$(NAME) $(MAIN_GO)

clean:
	go clean .
