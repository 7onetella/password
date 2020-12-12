#!/bin/sh -e
set +x

echo installing go-bindata-assetfs
go get -u github.com/go-bindata/go-bindata/...
go get -u github.com/elazarl/go-bindata-assetfs/...

echo generate go asset .go file
go-bindata-assetfs ui/...

go build ./...

echo installing gox
go get -u github.com/mitchellh/gox

echo cross compile
"${GOPATH}"/bin/gox -osarch="linux/amd64"

