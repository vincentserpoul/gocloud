#!/bin/bash

set -e -x

GOOS=linux GOARCH=amd64 go build -o testdeploy ./
docker build ./ --rm -t vincentserpoul/testdeploygolang -f ./Dockerfile