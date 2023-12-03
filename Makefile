# Variables
PROJECT_ROOT := $(shell pwd)
BINARY_PATH := $(PROJECT_ROOT)/build/user_account

database-migrate-up:
	go run ./cmd/command database migrate-up

.PHONY: build
build:
	go build -o build/user_account ./cmd/server/main.go