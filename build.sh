#!/bin/bash

cd ui

./build.sh

cd ../api

echo copying the emberjs asset from ui/dist/*
mkdir -p ui/
cp -r ../ui/dist/* ./ui/

./build.sh

# version=$(cat version | awk '{$1=$1;print}' | tail +1)

# docker build -t password .

# docker tag password:latest docker-registry.7onetella.net:5000/7onetella/password:"${version}"

# docker push docker-registry.7onetella.net:5000/7onetella/password:"${version}"

# vag docker deploy password-dev:"${version}"
