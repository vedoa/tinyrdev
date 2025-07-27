#!/bin/bash
set -e

mkdir -p dist

echo "Run test"

go test ./... -cover

VERSION=$(git describe --tags --abbrev=0 2>/dev/null || echo "dev")

echo "Building for Linux (amd64/arm64)"

GOOS=linux GOARCH=amd64 go build -ldflags="-X 'main.v=$VERSION'" -o dist/tinyrdev-"$VERSION"-linux-amd64
GOOS=linux GOARCH=arm64 go build -ldflags="-X 'main.v=$VERSION'" -o dist/tinyrdev-"$VERSION"-linux-arm64

echo "Building for macOS (arm64)"
GOOS=darwin GOARCH=arm64 go build -ldflags="-X 'main.v=$VERSION'" -o dist/tinyrdev-"$VERSION"-darwin-arm64

echo "Building for Windows (amd64)"
GOOS=windows GOARCH=amd64 go build -ldflags="-X 'main.v=$VERSION'" -o dist/tinyrdev-"$VERSION"-windows-amd64.exe

echo "All binaries built successfully."
