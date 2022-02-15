.PHONY: build
build:
	go build -v ./main
.DEFAULT_GOAL := build