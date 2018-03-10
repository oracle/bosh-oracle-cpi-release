#!/bin/bash

set -e 

pwd=`pwd`
export GOPATH=${pwd}/cpi-release-src
export PATH=$PATH:$GOPATH/bin

export CPITEST_CONFIG=${pwd}/oci-config/config
export CPITEST_PROFILE=DEFAULT

cd cpi-release-src/src/github.com/oracle/bosh-oracle-cpi
make testintci
