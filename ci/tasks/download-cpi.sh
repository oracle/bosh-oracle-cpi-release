#!/usr/bin/env bash

set -e

#Inputs
version=`cat dev-version-semver/number`

#Outputs
mkdir -p ${pwd}/candidate

# Create OCI config
echo "Creating oci config..."
OCI_DIR="$HOME/.oci"
OCI_API_KEY="$OCI_DIR/oci_api_key.pem"
OCI_CONFIG="$OCI_DIR/config"

mkdir -p ${OCI_DIR}
cat > ${OCI_API_KEY} <<EOF
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

chmod 600 ${OCI_API_KEY}
chmod 600 ${OCI_CONFIG}


cpi="bosh-oracle-cpi-dev-${version}.tgz"

# Download CPI
oci os object get -ns ${oracle_namespace} -bn ${oracle_bucket} --name ${cpi} --file ${cpi}
mv ${cpi} candidate/