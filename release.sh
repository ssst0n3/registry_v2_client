#!/bin/bash
set -x
# go get github.com/mitchellh/gox
cd $(dirname $(readlink -m $0))
mkdir -p bin/release
cd bin/release
CGO_ENABLED=0 gox -cgo=0 -osarch="linux/amd64" -osarch="linux/arm64" github.com/ssst0n3/registry_v2_client/cmd/regcli
