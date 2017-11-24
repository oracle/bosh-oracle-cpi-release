#!/usr/bin/env bash

set -e

deployment_dir="${PWD}/deployment"
cpi_release_name="bosh-oracle-cpi"
manifest_filename="director-manifest.yml"
state_filename="director-manifest-state.json"

echo "Setting up artifacts..."
cp ./candidate/*.tgz ${deployment_dir}/${cpi_release_name}.tgz
cp ./bosh-release/*.tgz ${deployment_dir}/bosh-release.tgz
cp ./stemcell/*.tgz ${deployment_dir}/stemcell.tgz
cp ./cpi-release-src/bosh-deployment/bosh.yml ${deployment_dir}/${manifest_filename}
cp ./cpi-release-src/bosh-deployment/cpi.yml ${deployment_dir}

pushd ${deployment_dir}
  function finish {
    echo "Final state of director deployment:"
    echo "=========================================="
    cat ${state_filename}
    echo "=========================================="

#    cp -r $HOME/.bosh_init ./
  }
  trap finish ERR

  echo "Using BOSH CLI version..."
  bosh -v

  ls -al 

  echo "Deploying BOSH Director..."
  bosh create-env --ops-file ./cpi.yml --vars-store ./creds.yml --state ${state_filename} --vars-file ./infra.yml ${manifest_filename}

  trap - ERR
  finish
popd
