#!/bin/sh -e

set -x

build() {

  # export DB_CONNSTR="postgres://dev:dev114@localhost/devdb"

  # export SERVER_ADDR=localhost:4242

  # export CRYPTO_TOKEN=test_crypto_token

  # go test -v ./...

  ./build.sh dev

  cd api

  # get gox for cross compilation
  go get -u github.com/mitchellh/gox

  # ${GOPATH}/bin/gox -osarch="linux/arm linux/amd64"
  ${GOPATH}/bin/gox -osarch="linux/amd64"

  # get jenkins server cpu architecture
  arch=$(lscpu | grep Archi | awk '{ print $2 }')
        
  case ${arch} in
  # raspberry pi 3 ARM cpu
  armv7l)
    # ${GOPATH}/bin/gox -osarch="linux/arm"
    # mv api_linux_arm password-api
    ;;
  # x86 compatible
  x86_64)
    # ${GOPATH}/bin/gox -osarch="linux/amd64"
    # mv api_linux_amd64 password-api
    ;;
  esac
  
  # chmod +x api_linux_arm
  chmod +x api_linux_amd64

  # mv api_linux_arm api_linux_arm_${BUILD_ID}    
  mv api_linux_amd64 api_linux_amd64_${BUILD_ID}    

} 

deploy() {
  cd api
  # upload to file server

  tar cvf api_dev_${BUILD_ID}.tar api_linux_amd64_${BUILD_ID} dev-*.pem

  scp -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no api_dev_${BUILD_ID}.tar pi@nas.7onetella.net:/mnt/uploads
  rm ./api_linux_amd64_${BUILD_ID}
  rm ./api_dev_${BUILD_ID}.tar

  # schedule a deployment
  cat ./password-api-dev.nomad | sed 's|BUILD_ID|'"${BUILD_ID}"'|g' > password-api-dev.nomad.${BUILD_ID}
  export NOMAD_ADDR=http://localhost:4646
  nomad job run ./password-api-dev.nomad.${BUILD_ID}
  rm ./password-api-dev.nomad.${BUILD_ID}
}

release() {
  cd api
  # upload to file server
  
  scp -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no ./api_linux_arm_${BUILD_ID} pi@nas.7onetella.net:/mnt/uploads
  rm ./api_linux_arm_${BUILD_ID}

  export NOMAD_ADDR=http://nomad.7onetella.net:4646
  # schedule a deployment
  nomad job run ./password-api-prod.nomad
}

case $1 in
build)
  build
  ;;
deploy)
  deploy 
  ;;
release)
  release 
  ;;
*)
  echo "Usage: ..."
  ;;
esac
