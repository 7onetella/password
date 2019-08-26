#!/bin/sh -e

build() {
    set -x
    go build ./...
    ls -l
    set +x
}

case $1 in
build)
  build
  ;;
*)
  echo "Usage: ..."
  ;;
esac