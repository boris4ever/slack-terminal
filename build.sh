#!/bin/bash


echo "Building for OSX"
env GOOS=darwin GOARCH=386 go build main.go
mv main slack-cli
zip slack-cli_OSX.zip slack-cli
rm slack-cli

echo "Building for Linux 64 BIT"
env GOOS=linux GOARCH=amd64 go build main.go
mv main slack-cli
zip slack-cli_Linux_64.zip slack-cli
rm slack-cli

echo "Building for Linux 32 BIT"
env GOOS=linux GOARCH=386 go build main.go
mv main slack-cli
zip slack-cli_Linux_32.zip slack-cli
rm slack-cli
