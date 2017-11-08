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

cp /usr/local/bin/bosh /usr/local/bin/bosh2
gem install bundle

echo "Setting up artifacts..."
cp ./stemcell/*.tgz stemcell.tgz

# get the terraform artifacts that were created previously
echo "Download oci infrastructure metadata..."
oci os object get -ns cloudfoundry -bn infra --name infra.tar.gz --file infra.tar.gz
tar zxof infra.tar.gz

echo "Download director credentials..."
oci os object get -ns cloudfoundry -bn infra --name creds.yml --file creds.yml

# path to the stemcell you want to use for testing
export BAT_STEMCELL="${PWD}/stemcell.tgz"

# path to the bat yaml file which is used to generate the deployment manifest
export BAT_DEPLOYMENT_SPEC="${PWD}/infra/bat.yml"

# BOSH CLI executable path
export BAT_BOSH_CLI=`which bosh2`

export BAT_DNS_HOST=8.8.8.8

# the name of infrastructure that is used by bosh deployment. Examples: aws, vsphere, openstack, warden.
export BAT_INFRASTRUCTURE=oci

# the type of networking being used: dynamic or manual.
export BAT_NETWORKING="manual"

# the path to ssh key, used by OS specs to ssh into BOSH VMs
export BAT_PRIVATE_KEY="$BOSH_SSH_KEY"

# Run tests with --fail-fast and skip cleanup in case of failure (optional)
export BAT_DEBUG_MODE=

export BOSH_ENVIRONMENT="$(bosh2 int infra/infra.yml --path /internal_ip)"
export BOSH_CLIENT=admin
export BOSH_CLIENT_SECRET="$(bosh2 int infra/infra.yml --path /admin_password)"
export BOSH_CA_CERT="$(bosh2 int creds.yml --path /director_ssl/ca)"

echo "Using BOSH CLI version..."
bosh2 -v

echo "Targeting BOSH director..."
bosh2 login

pushd bats
  echo "Installing gems..."
  bundle install

  echo "Running BOSH Acceptance Tests..."
  bundle exec rspec spec --tag ~vip_networking --tag ~dynamic_networking --tag ~root_partition --tag ~raw_ephemeral_storage --tag ~changing_static_ip --tag ~network_reconfiguration
popd
