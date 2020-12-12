#!/bin/sh
set +x

echo generate go asset .go file
go get -u github.com/go-bindata/go-bindata/...
go get -u github.com/elazarl/go-bindata-assetfs/...
go-bindata-assetfs ui/...

go build ./...

# get gox for cross compilation
go get -u github.com/mitchellh/gox

# ${GOPATH}/bin/gox -osarch="linux/arm linux/amd64"
echo cross compile
"${GOPATH}"/bin/gox -osarch="linux/amd64"

