#!/bin/sh

export DB_CONNSTR="postgres://dev:dev114@localhost/devdb"
export SERVER_ADDR=localhost:4242

refresh run -c ./config.yml
