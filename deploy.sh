#!/bin/bash

version=$(cat version | awk '{$1=$1;print}' | tail +1)

cd api

image=docker-registry.7onetella.net:5000/7onetella/password:"${version}"

docker build -t "${image}" .

docker push "${image}"

cd ..

vag docker deploy password-dev:"${version}"
