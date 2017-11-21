#!/usr/bin/env bash

set -e

#source bosh-cpi-src/ci/tasks/utils.sh
#source /etc/profile.d/chruby-with-ruby-2.1.2.sh

cpi_release_name="bosh-oracle-cpi"
semver=`cat version-semver/number`

image_path=dev_releases/${cpi_release_name}/${cpi_release_name}-${semver}.tgz
pushd bosh-cpi-src
  echo "Using BOSH CLI version..."
  bosh -v

  echo "Exposing release semver to bosh-oracle-cpi"
  echo ${semver} > "src/github.com/oracle/bosh-oracle-cpi/release"

  echo "Downloading dependency blobs (TEMPORARY)..."
  mkdir /tmp/blobs
  curl -o /tmp/blobs/go1.8.3.linux-amd64.tar.gz https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz
  bosh add-blob /tmp/blobs/go1.8.3.linux-amd64.tar.gz golang/go1.8.3.linux-amd64.tar.gz

  # We have to use the --force flag because we just added the src/github.com/oracle/bosh-oracle-cpi/release file
  echo "Creating CPI BOSH Release..."
  bosh create-release --name ${cpi_release_name} --version ${semver} --force --tarball="$image_path"
popd

echo -n $(sha1sum bosh-cpi-src/$image_path | awk '{print $1}') > bosh-cpi-src/$image_path.sha1

mv bosh-cpi-src/${image_path} candidate/
mv bosh-cpi-src/${image_path}.sha1 candidate/
