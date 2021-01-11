#!/bin/bash
# init registry
docker-compose up -d
# init test data
docker pull hello-world
docker tag hello-world 127.0.0.1:5000/dkdk/hello-world:v1
docker push 127.0.0.1:5000/dkdk/hello-world:v1