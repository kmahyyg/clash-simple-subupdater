#!/bin/bash
mkdir -p bin
rm -rf bin/clashsub*
CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -trimpath -ldflags '-s -w' -o bin/clashsub.amd64.ori ./cmd/main.go
strip bin/clashsub.amd64.ori
upx bin/clashsub.amd64.ori -o bin/clashsub.amd64
rm bin/clashsub.amd64.ori
