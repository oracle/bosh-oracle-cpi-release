#!/bin/bash

set -e

export GOPATH=$PWD/cpi-release-src
export PATH=$PATH:$GOPATH/bin

cd cpi-release-src/src/github.com/oracle/bosh-oracle-cpi
make test
