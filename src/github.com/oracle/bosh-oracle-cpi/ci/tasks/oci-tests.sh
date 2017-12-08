#!/bin/bash

set -e 

pwd=`pwd`
export GOPATH=${pwd}/cpi-src
export PATH=$PATH:$GOPATH/bin

export CPITEST_CONFIG=${pwd}/oci-config/config
export CPITEST_PROFILE=DEFAULT

cd cpi-src/src/github.com/oracle/bosh-oracle-cpi
make testintci
