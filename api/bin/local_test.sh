#!/bin/sh

DB_CONNSTR=$(consul kv get password-localhost-app/DB_CONNSTR)
export DB_CONNSTR

export SERVER_ADDR=localhost:4242
export INSECURE=true

go test -v ./...
