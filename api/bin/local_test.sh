#!/bin/sh

export DB_CONNSTR="postgres://dev:dev114@localhost/devdb"
export SERVER_ADDR=localhost:4242
export INSECURE=true

go test -v ./...
