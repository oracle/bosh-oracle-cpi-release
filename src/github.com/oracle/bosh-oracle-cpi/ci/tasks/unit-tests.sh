#!/bin/bash

set -e 

export GOPATH=$PWD/cpi-src
export PATH=$PATH:$GOPATH/bin

cd cpi-src/src/github.com/oracle/bosh-oracle-cpi
make test
