# Go example projects

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)

This repository contains a web app that able to save transaction and view them hourly

## Clone the project

```
$ git clone https://github.com/dsapoetra/btc.git
$ cd btc
```

## Run the project

Easiest way is to use docker-compose

```
$ cd btc
$ docker compose up
```

Current version of the project is running on http://localhost:8080/swagger/

This version doesn't have create db creation script, so you need to create the database manually with db name is db
migration is implemented in docker, will work after db created

## Run the test
```
go test -v ./...
```

## generate mock the test
```
make mock-gen
```