#!/bin/bash -e

set +x

build() {

  # execute in api project
  DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

  cd "$DIR" || exit

  cd "$DIR/ui"
  ./build.sh 

  cd "$DIR/api"
  ./build.sh

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

  echo uploading to file server
  tar czvf api_dev_${BUILD_ID}.tar.gz api_linux_amd64_${BUILD_ID} dev-*.pem

  scp -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no api_dev_${BUILD_ID}.tar.gz pi@nas.7onetella.net:/mnt/uploads
  rm ./api_linux_amd64_${BUILD_ID}
  rm ./api_dev_${BUILD_ID}.tar.gz

  echo scheduling deployment
  cat ./dev.nomad | sed 's|BUILD_ID|'"${BUILD_ID}"'|g' > dev.nomad.${BUILD_ID}
  export NOMAD_ADDR=http://localhost:4646
  nomad job run ./dev.nomad.${BUILD_ID}
  rm ./dev.nomad.${BUILD_ID}
}

run_test() {
  echo current location is $(pwd)
  export DB_CONNSTR="postgres://dev:dev114@localhost/devdb"
  export SERVER_ADDR=dev:4242
  export INSECURE=true
  go test -v ./...
}

release() {
  cd api
  # upload to file server
  
  scp -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no ./api_linux_arm_${BUILD_ID} pi@nas.7onetella.net:/mnt/uploads
  rm ./api_linux_arm_${BUILD_ID}

  export NOMAD_ADDR=http://nomad.7onetella.net:4646
  # schedule a deployment
  nomad job run ./prod.nomad
}

case $1 in
build)
  build
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
