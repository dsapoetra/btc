### Variables
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GOPATH ?= $(shell go env GOPATH)
GO111MODULE = on
CGO_ENABLED = 0


.PHONY: mockgen
mockgen:
	script/generate-mock.sh $(name)

