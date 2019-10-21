#!/bin/sh

export DB_CONNSTR="postgres://dev:dev114@localhost/devdb"
export SERVER_ADDR=localhost:4242
export CRYPTO_TOKEN=local_dev_crypto_token

refresh run -c ./config.yml
