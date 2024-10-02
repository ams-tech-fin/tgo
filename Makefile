include .env

APP_NAME = TGo
FILE_NAME = tgo

VERSION := $(shell cat VERSION)
GO_FILES := $(wildcard *.go)

MIGRATION_DIR=database/migrations

DB_DSN="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

.PHONY: all dev install container-up container-down build run clean version migrate-up migrate-down migrate-create

all: build

dev:
	go run github.com/cespare/reflex@latest -r '(\.go$$|\.env$$|VERSION$$|README.MD)' -s -- make run

install:
	go mod tidy

container-up:
	docker compose up -d

container-down:
	docker compose down

build:
	go build -o $(FILE_NAME) ./api/_cmd/app

run:
	go run ./api/_cmd/app/main.go

clean:
	rm -f $(FILE_NAME)

version:
	@echo "${APP_NAME} V$(VERSION)"

migrate-up:
	go run github.com/pressly/goose/v3/cmd/goose@latest -dir $(MIGRATION_DIR) postgres "$(DB_DSN)" up

migrate-down:
	go run github.com/pressly/goose/v3/cmd/goose@latest -dir $(MIGRATION_DIR) postgres "$(DB_DSN)" down

migrate-create:
	go run github.com/pressly/goose/v3/cmd/goose@latest create -dir $(MIGRATION_DIR) $(name) sql
