#!/bin/bash
set -x
docker-compose down
docker-compose up -d
docker pull ssst0n3/docker_secret:v1.0.0
docker tag ssst0n3/docker_secret:v1.0.0 127.0.0.1:5000/dkdk/hello-world:v1
docker push 127.0.0.1:5000/dkdk/hello-world:v1