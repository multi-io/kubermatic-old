#!/usr/bin/env bash

set -ex
export TAG=v2.3.1
export REPO=swaggerapi/swagger-codegen-cli

docker pull $REPO:$TAG
docker run --rm -v ${PWD}/..:/cmd $REPO:$TAG generate \
    -i /cmd/kubermatic-api/swagger.json \
    -l go -o /cmd/kubermatic-cli/client \
    -DpackageName=client
