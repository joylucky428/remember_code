#!/usr/bin/env bash
# 로컬 개발환경에 go 프로젝트를 빌드하고 container 실행

GOOS=linux go build -o ../../bin/remember_code_linux ../..

if [ $? -eq 0 ]; then
  docker-compose build
  docker-compose up
else
  echo "go build failed!"
fi