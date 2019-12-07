#!/bin/sh

export STAGE=$1

if [ "${1}" = "" ]; then
   echo
   echo "Usage: test.sh <stage>"
   echo
   echo "Example: test.sh localhost"
   echo "         test.sh dev"
   exit 1
fi

. bin/env.sh

SERVER_HOST=$(dig @127.0.0.1 -p 8600 password-dev-app.service.dc1.consul. SRV | grep "IN A" | awk '{ print $5 }')
SERVER_PORT=$(dig @127.0.0.1 -p 8600 password-dev-app.service.dc1.consul. SRV | grep "IN SRV" | awk '{ print $7 }')

if [ "${STAGE}" = "localhost" ]; then
   SERVER_HOST="localhost"
   SERVER_PORT=4242
fi

SERVER_ADDR="${SERVER_HOST}:${SERVER_PORT}"
export SERVER_ADDR
export INSECURE=true

go test -v ./...
