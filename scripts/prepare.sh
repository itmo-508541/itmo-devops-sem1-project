#!/bin/bash
go mod tidy
go mod download
go mod verify

go run cmd/main.go migrate
