#!/usr/bin/env bash

set -e

cpi_release_name="bosh-oracle-cpi"
semver=`cat final-version-semver/number`

pwd=`pwd`

#Inputs
src_repo=${pwd}/cpi-release-src
dev_release_tarball=${pwd}/candidate/*.tgz

#Outputs
mkdir -p ${pwd}/promoted

source ${src_repo}/ci/tasks/add-blobs.sh
cp -r ${src_repo} promoted/repo

release_artifact_path=${pwd}/promoted
tarball_name=${cpi_release_name}-${semver}.tgz
tarball_path=${release_artifact_path}/${tarball_name}
tarball_sha=${release_artifact_path}/${tarball_name}.sha1


pushd promoted/repo

  echo "Creating config/private.yml with blobstore secrets"
  cat > config/private.yml << EOF
---
blobstore:
  options:
    host: ${host}
    region: ${region_name}
    access_key_id: ${access_key_id}
    secret_access_key: ${secret_access_key}
    bucket_name: ${bucket}
    credentials_source: static
    signature_version: "4"
EOF

  echo "Using BOSH CLI version..."
  bosh -v

  # bosh add-blob
  addGolangBlobToRelease .

  echo "Creating BOSH Oracle CPI Final Release..."
  bosh finalize-release ${dev_release_tarball}  --name ${cpi_release_name} --version ${semver} --force

  cp ${dev_release_tarball} ${tarball_path}

  echo -n $(sha1sum $tarball_path | awk '{print $1}') > ${tarball_sha}
  echo

  echo "Built: ${tarball_path}"
  echo "sha1: " `cat ${tarball_sha}`

  git add .
  git config --global user.email bosh-build@oracle.com
  git config --global user.name CI
  git commit -m "BOSH Oracle CPI BOSH Release v${semver}"

popd

echo ${semver} > promoted/version
echo "BOSH Oracle CPI BOSH Release v${semver}" > promoted/annotation_message
