#!/bin/sh

gox -osarch="linux/arm"
scp api_linux_arm root@app1:/root/password
api_linux_arm
