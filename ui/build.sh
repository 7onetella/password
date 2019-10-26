#!/bin/sh
set -x

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$DIR/ui" || exit

npm install
rm -rf dist/*
ember build --environment=${1} 
