#!/usr/bin/env bash

CGO_ENABLED=0  go build -o simple-tool main.go

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o simple-tool_linux main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o simple-tool_windows main.go