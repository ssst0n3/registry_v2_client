#!/bin/bash
# init registry
ln -s $(pwd) /tmp/registry
./init.sh
#go test -count=1 -parallel 1 -v ./...
for s in $(go list ./...); do if ! go test -failfast -v -p 1 $s; then break; fi; done
