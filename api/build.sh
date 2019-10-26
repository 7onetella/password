#!/bin/sh
set +x

echo copying the emberjs asset from ui/dist/*
mkdir -p ui/
cp -r ../ui/dist/* ./ui/

# Jenkins path can be missing this
PATH=$PATH:.:~/bin

echo generate go asset .go file
go-bindata-assetfs ui/...

go build ./...

# get gox for cross compilation
go get -u github.com/mitchellh/gox

# ${GOPATH}/bin/gox -osarch="linux/arm linux/amd64"
echo cross compile
"${GOPATH}"/bin/gox -osarch="linux/amd64"

