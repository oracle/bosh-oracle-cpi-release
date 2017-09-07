#!/bin/bash
set -e

export BOSH_LOG_LEVEL=DEBUG
#export BOSH_LOG_LEVEL=INFO
export BOSH_LOG_PATH=./bosh.log

cp /dev/null $BOSH_LOG_PATH

bosh create-env --ops-file ./cpi.yml  --vars-store ./creds.yml --state director-state.json --vars-file $HOME/director-env-vars.yml bosh.yml
