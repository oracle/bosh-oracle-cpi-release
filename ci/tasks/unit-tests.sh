#!/usr/bin/env bash

set -e

export GOPATH=${PWD}/bosh-cpi-src
export PATH=${GOPATH}/bin:$PATH

cd ${PWD}/bosh-cpi-src/src/github.com/oracle/bosh-oracle-cpi
# TODO following make should run the test target
make build
