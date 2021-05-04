.PHONY: build clean install

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=resize11
MAINDIR=./cmd
BINARY_PATH=./build/$(BINARY_NAME)

build: clean
	$(GOBUILD) -o $(BINARY_PATH) -v $(MAINDIR)

clean:
	$(GOCLEAN)
	rm -rf ./build
	mkdir ./build

install: build
	cp -f $(BINARY_PATH) $(GOPATH)/bin
