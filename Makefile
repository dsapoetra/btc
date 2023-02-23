### Variables
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GOPATH ?= $(shell go env GOPATH)
GO111MODULE = on
CGO_ENABLED = 0


.PHONY: mockgen
mockgen:
	script/generate-mock.sh $(name)

.PHONY: create-migration
create-migration :
	goose -dir migrations create $(input) sql

.PHONY: all-db-migrate
all-db-migrate:
	goose -dir migrations/ postgres "user=postgres password=D54poetra dbname=btc sslmode=disable" up

