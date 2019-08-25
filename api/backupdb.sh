#!/bin/sh

echo initiating database backup
ssh root@app1 '/root/backupdb.sh'

echo copying the backup to local
scp root@app1:/root/password.sql .

mv password.sql ~/gdrive/backups 
