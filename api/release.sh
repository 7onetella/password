#!/bin/sh

echo cross compile
"${GOPATH}"/bin/gox -osarch="linux/amd64"

echo stop password service
ssh root@app1 'systemctl stop password'

echo copying password to remote server
scp api_linux_arm root@app1:/root/password

echo starting password service
ssh root@app1 'systemctl start password'
rm api_linux_arm
