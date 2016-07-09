#!/bin/bash

version=$1

# windows x86
GOOS=windows GOARCH=amd64 go build -o _build/simplehttpserver.exe -ldflags "-X main.version=$version" simplehttpserver.go
zip -j9 _build/win.x64_$version.zip _build/simplehttpserver.exe
rm _build/simplehttpserver.exe

# linux x86
GOOS=linux GOARCH=amd64 go build -o _build/simplehttpserver -ldflags "-X main.version=$version" simplehttpserver.go
zip -j9 _build/linux.x64_$version.zip _build/simplehttpserver
rm _build/simplehttpserver

# darwin x86
GOOS=darwin GOARCH=amd64 go build -o _build/simplehttpserver -ldflags "-X main.version=$version" simplehttpserver.go
zip -j9 _build/darwin.x64_$version.zip _build/simplehttpserver
rm _build/simplehttpserver
