#!/bin/bash -e

set +x

build() {

  arch=$1

  # execute in api project
  DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

  cd "$DIR" || exit

  cd "$DIR/ui"
  ./build.sh 

  cd "$DIR/api"
  ./build.sh

  case ${arch} in
  # raspberry pi 3 ARM cpu
  armv7l)
     ${GOPATH}/bin/gox -osarch="linux/arm"
     mv api_linux_arm password
    ;;
  # x86 compatible
  linux)
     ${GOPATH}/bin/gox -osarch="linux/amd64"
     mv api_linux_amd64 password
    ;;
  # Mac OSX
  darwin)
     ${GOPATH}/bin/gox -osarch="darwin/amd64"
     mv api_linux_amd64 password
    ;;
  esac
  
  # chmod +x api_linux_arm
  chmod +x password

}

deploy() {
  echo deploying...
}

run_test() {
  echo running tests...
}

release() {
  echo reeleasing....
}

case $1 in
build)
  build $2
  ;;
deploy)
  deploy 
  ;;
test)
  run_test
  ;;
release)
  release 
  ;;
*)
  echo "Usage: ..."
  ;;
esac
