#!/usr/bin/env bash

set -e

source bosh-cpi-src/ci/tasks/utils.sh

check_param oracle_tenancy
check_param oracle_user
check_param oracle_compartment_id
check_param oracle_compartment_name
check_param oracle_region
check_param oracle_ad
check_param oracle_fingerprint
check_param oracle_apikey
check_param oracle_ssh_public_key

echo "Creating oci config..."
OCI_DIR="$HOME/.oci"
OCI_API_KEY="$OCI_DIR/oci_api_key.pem"
OCI_CONFIG="$OCI_DIR/config"

mkdir -p $OCI_DIR
cat > $OCI_API_KEY <<EOF
${oracle_apikey}
EOF

cat > $OCI_CONFIG <<EOF
[DEFAULT]
user=${oracle_user}
tenancy=${oracle_tenancy}
region=${oracle_region}
key_file=$OCI_API_KEY
fingerprint=${oracle_fingerprint}
EOF

chmod 600 $OCI_API_KEY

deployment_dir="${PWD}/deployment"
cpi_release_name=bosh-oracle-cpi
manifest_filename="director-manifest.yml"
state_filename="director-manifest-state.json"

echo "Setting up artifacts..."
#cp ./bosh-cpi-release/*.tgz ${deployment_dir}/${cpi_release_name}.tgz
cp ./bosh-release/*.tgz ${deployment_dir}/bosh-release.tgz
cp ./stemcell/*.tgz ${deployment_dir}/stemcell.tgz
cp ./bosh-cpi-src/bosh-deployment/bosh.yml ${deployment_dir}/${manifest_filename}
cp ./bosh-cpi-src/bosh-deployment/cpi.yml ${deployment_dir}

oci os object get -ns cloudfoundry -bn infra --name infra.tar.gz --file infra.tar.gz
tar zxof infra.tar.gz
cp ./infra/infra.yml ${deployment_dir}

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

  oci os object put -ns cloudfoundry -bn infra --force --name creds.yml --file creds.yml

  ls -al 

  trap - ERR
  finish
popd
