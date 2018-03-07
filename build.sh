#!/usr/bin/env bash

rm -rf bin &> /dev/null
mkdir bin &> /dev/null

echo "Building Darwin"
GOOS=darwin  GOARCH=amd64 go build -o bin/tagger-darwin-amd64
shasum -a 256 < bin/tagger-darwin-amd64
echo "Building Linux"
GOOS=linux   GOARCH=amd64 go build -o bin/tagger-linux-amd64
echo "Building Windows"
GOOS=windows GOARCH=amd64 go build -o bin/tagger-windows-amd64.exe
