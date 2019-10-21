#!/bin/sh

export DB_CONNSTR="postgres://dev:dev114@localhost/devdb"
export SERVER_ADDR=localhost:4242

go test -v ./...
