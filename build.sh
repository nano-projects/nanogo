#!/bin/bash

rm -rf bin && mkdir bin && cd bin
echo "clean and make bin"

echo "build nanogo (mac x64) ..."
GOOS=darwin GOARCH=amd64 go build ../. && tar -zcvf nanogo_darwin_amd64.tar.gz nanogo && rm -rf nanogo

echo "build nanogo (windows x86) ..."
GOOS=windows GOARCH=386 go build ../. && tar -zcvf nanogo_windows_386.tar.gz nanogo.exe && rm -rf nanogo.exe

echo "build nanogo (windows x64) ..."
GOOS=windows GOARCH=amd64 go build ../. && tar -zcvf nanogo_windows_amd64.tar.gz nanogo.exe && rm -rf nanogo.exe

echo "build nanogo (linux x86) ..."
GOOS=linux GOARCH=386 go build ../. && tar -zcvf nanogo_linux_386.tar.gz nanogo && rm -rf nanogo

echo "build nanogo (linux x64) ..."
GOOS=linux GOARCH=amd64 go build ../. && tar -zcvf nanogo_linux_amd64.tar.gz nanogo && rm -rf nanogo

echo "build all arch Successfully"
