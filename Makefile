.PHONY: build
build:
# this build us app to exe
	go build -v ./cmd/apiserver

.DEFAULT_GOAL := build
#   For win run
# > make.exe