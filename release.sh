#!/usr/bin/bash

set -e

OS_LIST="linux darwin windows"
for OS in $OS_LIST; do
    echo "Building for $OS"
    TARGET=bin/auto-mail-${OS}-amd64
    if [ $OS == "windows" ]; then
        TARGET=bin/auto-mail-${OS}-amd64.exe
    fi
    GOOS=$OS GOARCH=amd64 go build -o $TARGET cmd/main.go
done