#!/usr/bin/env bash

#mac 64
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./dist/mac/simple-tool main.go
#linux 64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./dist/linux/simple-tool main.go
#windows 64
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./dist/windows/simple-tool main.go