#!/bin/bash
CURRENT_ARCH=$1
mkdir -p bin
rm -rf bin/clashsub-"${CURRENT_ARCH}"
CGO_ENABLED=0 GOOS=linux GOARCH="${CURRENT_ARCH}" GOMIPS=softfloat go build -trimpath -ldflags '-s -w' -o bin/clashsub."${CURRENT_ARCH}".ori ./cmd/main.go
strip bin/clashsub."${CURRENT_ARCH}".ori
upx bin/clashsub."${CURRENT_ARCH}".ori -o bin/clashsub."${CURRENT_ARCH}"
rm bin/clashsub."${CURRENT_ARCH}".ori
