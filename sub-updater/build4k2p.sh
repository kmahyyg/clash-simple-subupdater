#!/bin/bash
CGO_ENABLED=0 GOARCH=mipsle GOOS=linux GOMIPS=softfloat go build -trimpath -ldflags '-s -w' -o clashsub.ori ./cmd/main.go
strip ./clashsub.ori
upx ./clashsub.ori -o ./clashsub
rm ./clashsub.ori
