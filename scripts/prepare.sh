#!/bin/bash
go mod tidy
go mod download
go mod verify

mkdir --parent ./bin
go build -o ./bin/app ./cmd/main.go

./bin/app migrate
