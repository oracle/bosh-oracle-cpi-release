#!/usr/bin/env bash

set -e

cpi_release_name="bosh-oracle-cpi"
semver=`cat version-semver/number`
golang_ver=1.8.3

pwd=`pwd`
tarball_name=${cpi_release_name}-dev-${semver}.tgz
tarball_path=${pwd}/artifacts/${tarball_name}
tarball_sha=${pwd}/artifacts/${tarball_name}.sha1
release_dir=${pwd}/cpi-release

echo "Using BOSH CLI version..."
bosh -v

echo "Populating blobstore"
golang_distro=`basename /blobstore/golang/go*.gz`
bosh add-blob --dir ${release_dir} /blobstore/golang/${golang_distro} golang/${golang_distro}

echo "Creating Oracle CPI BOSH Dev Release..."
bosh create-release --dir ${release_dir} --name ${cpi_release_name} --version ${semver} --force --tarball="$tarball_path"

echo -n $(sha1sum $tarball_path | awk '{print $1}') > ${tarball_sha} 

echo "Built: ${tarball_path}"
echo "sha1: " `cat ${tarball_sha}` 
