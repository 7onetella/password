#!/bin/sh
set -x

cd ui

rm -rf dist/* && ember build --environment=${1} && cp -r dist/* ../api/ui/

cd ..

cd api

echo current location is $(pwd)

go get github.com/jteeuwen/go-bindata/...
go get github.com/elazarl/go-bindata-assetfs/...

echo PATH = ${PATH}

which go-bindata-assetfs

go-bindata-assetfs ui/...

go build ./...

cd ..