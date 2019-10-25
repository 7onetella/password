#!/bin/sh

export DB_CONNSTR="postgres://dev:dev114@keepass/devdb"
export SERVER_ADDR=keepass:443
export CRYPTO_TOKEN=production_crypto_token

refresh run -c ./config.yml
