#!/usr/bin/env bash

set -e

source bosh-cpi-src/ci/tasks/utils.sh

deployment_dir="${PWD}/deployment"
manifest_filename="director-manifest.yml"
state_filename="director-manifest-state.json"

pushd ${deployment_dir}

  echo "Using BOSH CLI version..."
  bosh -v

  echo "Deleting BOSH Director..."
  bosh delete-env --ops-file ./cpi.yml --vars-store ./creds.yml --state ${state_filename} --vars-file ./infra.yml ${manifest_filename}
popd
