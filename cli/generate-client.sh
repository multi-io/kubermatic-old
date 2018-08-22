#!/usr/bin/env bash

set -ex
export TAG=v2.3.1
export REPO=swaggerapi/swagger-codegen-cli

cp ${PWD}/../api/cmd/kubermatic-api/swagger.json .

docker pull $REPO:$TAG
docker run --rm -v ${PWD}:/gen $REPO:$TAG generate \
    -i /gen/swagger.json \
    -l go -o /gen/api \
    -DpackageName=api

rm swagger.json