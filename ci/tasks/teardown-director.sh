#!/usr/bin/env bash

set -e

deployment_dir="${PWD}/deployment"
manifest_filename="director-manifest.yml"
state_filename="director-state.json"

pushd ${deployment_dir}

  echo "Using BOSH CLI version..."
  bosh -v

  echo "Deleting BOSH Director..."
  bosh delete-env --ops-file ./cpi.yml --ops-file ./local.yml --vars-store ./creds.yml --state ${state_filename} --vars-file ./director-env-vars.yml ${manifest_filename}
popd
