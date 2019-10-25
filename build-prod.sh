#!/bin/sh

cd ui

rm -rf dist/* && ember build --environment=production && cp -r dist/* ../api/ui/

cd ../api

go-bindata-assetfs ui/...

cd ..