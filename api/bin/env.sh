#!/bin/sh

DB_CONNSTR=$(consul kv get password-"${STAGE}"-app/DB_CONNSTR)
export DB_CONNSTR

CRYPTO_TOKEN=$(consul kv get password-"${STAGE}"-app/CRYPTO_TOKEN)
export CRYPTO_TOKEN

CREDENTIAL=$(consul kv get password-"${STAGE}"-app/CREDENTIAL)
export CREDENTIAL