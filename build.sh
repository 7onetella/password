#!/bin/sh
set -x

echo $GOPATH
export GOPATH=~

cd ui

rm -rf dist/* && ember build --environment=${1} && cp -r dist/* ../api/ui/

cd ../api

~/bin/go-bindata-assetfs ui/...

cd ..