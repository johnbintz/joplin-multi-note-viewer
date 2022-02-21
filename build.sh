#!/bin/bash

VERSION=$(cat VERSION)

rm -Rf build
mkdir build
cd client
nvm use
npm run build
cd ../server
rm -Rf dist
mv ../client/dist .


GOOS=linux GOARCH=amd64 go build -o ../build/viewer-${VERSION}-linux-amd64 .
GOOS=linux GOARCH=arm64 go build -o ../build/viewer-${VERSION}-linux-arm64 .
GOOS=windows GOARCH=amd64 go build -o ../build/viewer-${VERSION}-windows-amd64 .
GOOS=windows GOARCH=arm64 go build -o ../build/viewer-${VERSION}-windows-arm64 .
GOOS=darwin GOARCH=amd64 go build -o ../build/viewer-${VERSION}-darwin-amd64 .
GOOS=darwin GOARCH=arm64 go build -o ../build/viewer-${VERSION}-darwin-arm4 .

cd ../build
cp ../config.yml.example ./config.yml

echo "Done. App is in build/"
echo "Have fun."
