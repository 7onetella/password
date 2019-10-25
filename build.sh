#!/bin/sh
set -x

cd ui

rm -rf dist/* && ember build --environment=${1} && cp -r dist/* ../api/ui/

cd ..

cd api

echo current location is $(pwd)

~/bin/go-bindata-assetfs ui/...

go build ./...

cd ..