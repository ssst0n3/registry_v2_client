#!/bin/bash
# init registry
docker-compose down
docker-compose up -d
# init test data
docker pull hello-world
docker tag hello-world 127.0.0.1:5000/dkdk/hello-world:v1
docker push 127.0.0.1:5000/dkdk/hello-world:v1
go test -count=1 -parallel 1 -v ./...