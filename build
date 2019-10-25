#!/bin/sh
set -x

cd ui

rm -rf dist/* && ember build --environment=${1} && cp -r dist/* ../api/ui/

cd ../api

go-bindata-assetfs ui/...

cd ..