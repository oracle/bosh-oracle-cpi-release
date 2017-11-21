#!/usr/bin/env bash

set -e

source bosh-cpi-src/ci/tasks/utils.sh

#check_param release_blobs_access_key
#check_param release_blobs_secret_key

# Version info
semver_version=`cat release-version-semver/number`
echo $semver_version > promoted/semver_version
echo "BOSH Oracle CPI BOSH Release v${semver_version}" > promoted/annotation_message

today=$(date +%Y-%m-%d)
cp -r bosh-cpi-src promoted/repo

# CPI vars
cpi_release_name="bosh-oracle-cpi"
cpi_blob=${cpi_release_name}-${semver_version}.tgz
cpi_link=https://storage.googleapis.com/bosh-cpi-artifacts/bosh-oracle-cpi-$semver_version.tgz

dev_release=$(echo $PWD/bosh-cpi-release/*.tgz)

pushd promoted/repo
  echo "Creating config/private.yml with blobstore secrets"
  set +x

  echo "Using BOSH CLI version..."
  bosh -v

  echo "Finalizing CPI BOSH Release..."
  bosh finalize release ${dev_release} --version ${semver_version}

  rm config/private.yml

  # Insert CPI details into README.md
  # Template markers in the README
  cpi_marker="\[//\]: # (new-cpi)"
  cpi_sha=$(sha1sum releases/$cpi_release_name/$cpi_blob | awk '{print $1}')
  new_cpi="|[$semver_version]($cpi_link)|$cpi_sha|$today|"
  sed -i "s^$cpi_marker^$new_cpi\n$cpi_marker^" README.md

  git diff | cat
  git add .

  git config --global user.email cf-bosh-eng@pivotal.io
  git config --global user.name CI
  git commit -m "BOSH Oracle CPI BOSH Release v${semver_version}"

  mv releases/$cpi_release_name/$cpi_blob ../
  echo $cpi_sha > ../$cpi_blob.sha1
popd

