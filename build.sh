#!/bin/bash

rm -rf bin && mkdir bin && cd bin && mkdir -p darwin/amd64 windows/386 windows/amd64 linux/386 linux/amd64
echo "clean and make bin"

echo "build nanogo (mac x64) ..."
GOOS=darwin GOARCH=amd64 go build ../. && mv nanogo darwin/amd64

echo "build nanogo (windows x86) ..."
GOOS=windows GOARCH=386 go build ../. && mv nanogo.exe windows/386

echo "build nanogo (windows x64) ..."
GOOS=windows GOARCH=amd64 go build ../. && mv nanogo.exe windows/amd64

echo "build nanogo (linux x86) ..."
GOOS=linux GOARCH=386 go build ../. && mv nanogo linux/386

echo "build nanogo (linux x64) ..."
GOOS=linux GOARCH=amd64 go build ../. && mv nanogo linux/amd64

echo "build all arch Successfully"
