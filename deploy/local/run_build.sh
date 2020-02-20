#!/usr/bin/env bash

# build bin for linux
GOOS=linux go build -o ../../bin/remember_code_linux ../..

# push images
docker push joylucky/remember_code:1.0
docker push joylucky/remember_code_nginx:1.0