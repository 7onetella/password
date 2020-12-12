#!/bin/bash

docker run -d --name password -p 8080:8080 \
-e STAGE=devpass \
-e HTTP_PORT=8080 \
-e DB_CONNSTR=postgres://dev:dev114@tmt-vm11.7onetella.net:5432/devdb \
-e CRYPTO_TOKEN=test \
-e CREDENTIAL="John.Smith@example.com:users91234" \
password 
