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

BOSH_SSH_KEY="${PWD}/bosh-ssh.key"
cat > $BOSH_SSH_KEY <<EOF
${oracle_ssh_private_key}
EOF
chmod 600 $BOSH_SSH_KEY

# Setup Go and run tests
export GOPATH=${PWD}/bosh-cpi-src
export PATH=${GOPATH}/bin:$PATH

#check_go_version $GOPATH

cd ${PWD}/bosh-cpi-src/src/github.com/oracle/bosh-oracle-cpi
env
make testintci
