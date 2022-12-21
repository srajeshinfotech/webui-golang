.DEFAULT_GOAL := all

sidecar_root_dir := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))..
UNAME := $(shell uname -m)

prebuild:
    $(eval PRESENT_GOPATH := `go env GOPATH`)

compile: prebuild
	echo "Compiling for every OS and Platform"
	echo "GOPATH is $(PRESENT_GOPATH)"
	rm -rf bin/*
	#GOPATH=$(PRESENT_GOPATH) PKG_CONFIG_PATH=/usr/local/lib64/pkgconfig/ GOOS=linux GOARCH=amd64 go build -v -o bin/ui src/main.go
	GOPATH=$(PRESENT_GOPATH) PKG_CONFIG_PATH=/usr/local/lib64/pkgconfig/ GOOS=darwin GOARCH=amd64 go build -v -o bin/ui src/main.go

copyover:
	echo "Copying over"
	cp -rfvp static bin/

all: compile copyover

clean:
	rm -rf src/github.com
	rm -rf bin

