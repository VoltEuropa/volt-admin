#!/bin/sh
rm -rf bin
mkdir bin

GOOS=darwin go build -o bin/darwin/volt-admin
GOOS=linux go build -o bin/linux/volt-admin
GOOS=windows GOARCH=amd64 go build -o bin/windows/volt-admin.exe

SHA=$(git rev-parse --short HEAD)
tar -czvf bin/macos-volt-admin-$SHA.tar.gz -C bin/darwin/ volt-admin
tar -czvf bin/linux-volt-admin-$SHA.tar.gz -C bin/linux/ volt-admin
tar -czvf bin/windows-volt-admin-$SHA.tar.gz -C bin/windows/ volt-admin.exe

rm -rf bin/darwin
rm -rf bin/linux
rm -rf bin/windows
