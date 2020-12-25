#!/bin/bash
mkdir -p bin
rm -rf bin/clashsub*
CGO_ENABLED=0 GOARCH=mipsle GOOS=linux GOMIPS=softfloat go build -trimpath -ldflags '-s -w' -o bin/clashsub.mips24kcle.ori ./cmd/main.go
strip bin/clashsub.mips24kcle.ori
upx bin/clashsub.mips24kcle.ori -o bin/clashsub.mips24kcle
rm bin/clashsub.mips24kcle.ori
