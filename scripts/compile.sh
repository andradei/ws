#! /bin/bash

set -e

echo '-- Compiling for UNIX'

GOOS=linux GOARCH=amd64 go build -o bin/ws_linux

echo '-- Compiling for macOS'

GOOS=darwin GOARCH=amd64 go build -o bin/ws_darwin