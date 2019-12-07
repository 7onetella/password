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

export SERVER_ADDR=password-${STAGE}-app.7onetella.net:4242
export INSECURE=true

go test -v ./...
