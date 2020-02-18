#!/usr/bin/env bash

GOOS=linux go build -o ./remember_code_linux

if [ $? -eq 0 ]; then
  docker-compose build
  docker-compose up
else
  echo "go build failed!"
fi