#!/bin/sh
set -x

# execute in api project
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$DIR" || exit

mkdir -p ui/
cp -r ../ui/dist/* ./ui/

# Jenkins path can be missing this
PATH=$PATH:.:~/bin

go-bindata-assetfs ui/...

go build ./...

# get gox for cross compilation
go get -u github.com/mitchellh/gox

# ${GOPATH}/bin/gox -osarch="linux/arm linux/amd64"
${GOPATH}/bin/gox -osarch="linux/amd64"

